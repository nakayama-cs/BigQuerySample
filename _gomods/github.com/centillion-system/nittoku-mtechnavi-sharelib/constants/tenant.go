package constants

const (
	// グローバル共有レコード用のテナントID
	// システム内での重複を避けるために、UUIDバージョン1から5の使用は避ける
	// 上記から16進表記でxxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx（Mの桁がバージョン（1から5））となっているバージョンを0とすることで他テナントIDとの衝突を回避する
	GlobalTenantId = "efd7c12d-372b-0ba4-bfdb-06faa9303edf"

	// 運営テナントID
	// グローバル共有レコード用のテナントIDと同じく、バージョンを0とする
	AdminTenantId = "b80a703d-d14c-0b7a-a0e8-0fed00cb7723"
)
