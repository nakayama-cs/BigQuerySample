// See https://docs.google.com/spreadsheets/d/1RtI6DkEVRbeg-kOnEmGfn1RIpful0R8JTksZzYor4m8/edit#gid=251201936
package message

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

var (
	fallbackLanguage = language.Japanese

	bundles map[string]*i18n.Bundle

	bundlesMu sync.Mutex

	// intended to overwrite
	Loader LoaderFunc = func(context.Context, language.Tag) (map[string]string, error) {
		return nil, nil
	}
)

type LoaderFunc func(context.Context, language.Tag) (map[string]string, error)

type TranslateFunc func(string, ...any) string

func T(lang string) TranslateFunc {
	langTag, err := language.Parse(lang)
	if err != nil {
		logrus.WithError(err).Warnf("invalid language: %q", lang)
		langTag = fallbackLanguage
	}
	return func(messageName string, args ...any) string {
		bundle := getBundle(langTag)
		localizer := i18n.NewLocalizer(bundle)
		s, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: messageName,
		})
		if err != nil {
			logrus.WithError(err).Warnf("message %q is not found", messageName)
			l := make([]string, 0, 1+len(args))
			l = append(l, fmt.Sprintf("%q", messageName))
			for _, arg := range args {
				l = append(l, fmt.Sprint(arg))
			}
			return "[" + strings.Join(l, ", ") + "]"
		}
		return renderString(s, args...)
	}
}

func getBundle(lang language.Tag) *i18n.Bundle {
	bundlesMu.Lock()
	defer bundlesMu.Unlock()

	base, _, _ := lang.Raw()
	out, ok := bundles[base.String()]
	if ok {
		return out
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dict, err := Loader(ctx, lang)
	if err != nil {
		logrus.WithError(err).Warnf("failed to load dictionary for language: %v", lang.String())
	}
	out = i18n.NewBundle(lang)
	for k, v := range dict {
		out.MustAddMessages(lang, &i18n.Message{
			ID:    k,
			Other: v,
		})
	}
	if bundles == nil {
		bundles = map[string]*i18n.Bundle{}
	}
	bundles[base.String()] = out
	return out
}

func renderString(s string, args ...any) string {
	if len(args) == 0 {
		return s
	}

	// XXX: 変数展開は、単純な文字列置換で実現する
	for i, arg := range args {
		s = strings.ReplaceAll(s, fmt.Sprintf("{$%d}", i+1), fmt.Sprint(arg))
	}
	return s
}
