package littlecow

import (
	"log/slog"
	"path/filepath"
)

func Main() int {
	slog.Debug("littlecow", "test", true)

	return 0
}

func setPartialPath(source *slog.Source) {
	fileName := filepath.Base(source.File)
	parentDir := filepath.Base(filepath.Dir(source.File))

	source.File = filepath.Join(parentDir, fileName)
}

func Replace(groups []string, attr slog.Attr) slog.Attr {
	if attr.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	if attr.Key == slog.SourceKey {
		source, _ := attr.Value.Any().(*slog.Source)
		if source != nil {
			setPartialPath(source)
		}
	}
	return attr
}
