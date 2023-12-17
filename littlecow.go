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

func removeTimestamp(groups []string, attr *slog.Attr) {
	if attr.Key == slog.TimeKey && len(groups) == 0 {
		*attr = slog.Attr{}
	}
}

func adjustSourcePath(groups []string, attr *slog.Attr) {
	if attr.Key == slog.SourceKey {
		source, _ := attr.Value.Any().(*slog.Source)
		if source != nil {
			setPartialPath(source)
		}
	}
}

func Replace(groups []string, attr slog.Attr) slog.Attr {
	removeTimestamp(groups, &attr)
	adjustSourcePath(groups, &attr)
	return attr
}

func Replace2(groups []string, attr slog.Attr) slog.Attr {
	adjustSourcePath(groups, &attr)
	return attr
}
