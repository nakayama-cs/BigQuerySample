/*
 * このパッケージは、DataProxy APIを利用するためのコーデックを提供する。
 *
 * DataProxyでは、利用元の *.proto ファイルに依存せず、Protocol Buffersの仕様を殺さないような処理を行っている。
 * そのため、DataProxyの利用にあたっては、利用元の責務として *.proto ファイル定義を参照しつつ、DataProxy専用の規則に従ってMessageを変換する必要がある。
 * このパッケージはDataProxyの利用元で実行されることを前提としており、上述の変換処理を実装する。
 */
package codec

import (
	codecinternal "mtechnavi/sharelib/dataproxy/codec/internal"
)

const (
	pkgName = "dataproxy/codec"
)

func GetPrimaryKeyName(typeName string) string {
	return codecinternal.GetPrimaryKeyName(typeName)
}
