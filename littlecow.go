package littlecow

import (
	"log/slog"
	"os"
	"path/filepath"
)

func Main() int {
	opts := OpinionatedHandlerOptions(slog.LevelDebug, Unmodified)
	handler := slog.NewTextHandler(os.Stderr, opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	slog.Debug("debug", "debug", "debug")
	slog.Error("error", "error", "error")

	return 0
}

type ReplaceFunc func(groups []string, attr slog.Attr) slog.Attr

func setPartialPath(source *slog.Source) {
	fileName := filepath.Base(source.File)
	parentDir := filepath.Base(filepath.Dir(source.File))

	source.File = filepath.Join(parentDir, fileName)
}

func _removeTimestamp(groups []string, attr *slog.Attr) {
	if attr.Key == slog.TimeKey && len(groups) == 0 {
		*attr = slog.Attr{}
	}
}

func _truncateSourcePath(groups []string, attr *slog.Attr) {
	if attr.Key == slog.SourceKey {
		source, _ := attr.Value.Any().(*slog.Source)
		if source != nil {
			setPartialPath(source)
		}
	}
}

func RemoveTimestampAndTruncateSource(groups []string, attr slog.Attr) slog.Attr {
	_removeTimestamp(groups, &attr)
	_truncateSourcePath(groups, &attr)
	return attr
}

func RemoveTimestamp(groups []string, attr slog.Attr) slog.Attr {
	_removeTimestamp(groups, &attr)
	return attr
}

func TruncateSourcePath(groups []string, attr slog.Attr) slog.Attr {
	_truncateSourcePath(groups, &attr)
	return attr
}

func Unmodified(groups []string, attr slog.Attr) slog.Attr {
	return attr
}

func OpinionatedHandlerOptions(level slog.Level, replaceFunc ReplaceFunc) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource:   true,
		Level:       level,
		ReplaceAttr: replaceFunc,
	}
}
