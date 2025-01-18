package resources

import (
	_ "embed" // Пакет для встраивания файлов
)

//go:embed css/styles.css
var CssContent string
