package flags

import (
	"fmt"

	"github.com/tkuchiki/alp/options"

	"github.com/tkuchiki/alp/stats"
	"gopkg.in/alecthomas/kingpin.v2"
)

type GlobalFlags struct {
	Config         string
	File           string
	Dump           string
	Load           string
	Sort           string
	Reverse        bool
	QueryString    bool
	Format         string
	NoHeaders      bool
	Limit          int
	Location       string
	Output         string
	MatchingGroups string
	Filters        string
}

var SortKeys = []string{
	"max",
	"min",
	"avg",
	"sum",
	"count",
	"uri",
	"method",
	"max-body",
	"min-body",
	"avg-body",
	"sum-body",
	"p1",
	"p50",
	"p99",
	"stddev",
}

var Formats = []string{
	"table",
	"tsv",
}

var SortOptions = map[string]string{
	"max":      stats.SortMaxResponseTime,
	"min":      stats.SortMinResponseTime,
	"avg":      stats.SortAvgResponseTime,
	"sum":      stats.SortSumResponseTime,
	"count":    stats.SortCount,
	"uri":      stats.SortUri,
	"method":   stats.SortMethod,
	"max-body": stats.SortMaxResponseBodyBytes,
	"min-body": stats.SortMinResponseBodyBytes,
	"avg-body": stats.SortAvgResponseBodyBytes,
	"sum-body": stats.SortSumResponseBodyBytes,
	"p1":       stats.SortP1ResponseTime,
	"p50":      stats.SortP50ResponseTime,
	"p99":      stats.SortP99ResponseTime,
	"stddev":   stats.SortStddevResponseTime,
}

func NewGlobalFlags() *GlobalFlags {
	return &GlobalFlags{}
}

func (f *GlobalFlags) InitGlobalFlags(app *kingpin.Application) {
	app.Flag("config", "The configuration file").
		Short('c').StringVar(&f.Config)
	app.Flag("file", "The access log file").
		StringVar(&f.File)
	app.Flag("dump", "Dump profiled data as YAML").
		Short('d').StringVar(&f.Dump)
	app.Flag("load", "Load the profiled YAML data").
		Short('l').StringVar(&f.Load)
	app.Flag("sort", "Output the results in sorted order").
		Default(options.DefaultSortOption).EnumVar(&f.Sort, SortKeys...)
	app.Flag("reverse", "Sort results in reverse order").
		Short('r').BoolVar(&f.Reverse)
	app.Flag("query-string", "Include the URI query string.").
		Short('q').BoolVar(&f.QueryString)
	app.Flag("format", "The output format (table or tsv)").
		Default(options.DefaultFormatOption).EnumVar(&f.Format, Formats...)
	app.Flag("noheaders", "Output no header line at all (only --format=tsv)").
		BoolVar(&f.NoHeaders)
	app.Flag("limit", "The maximum number of results to display.").
		Default(fmt.Sprint(options.DefaultLimitOption)).IntVar(&f.Limit)
	app.Flag("location", "Location name for the timezone").
		Default(options.DefaultLocationOption).StringVar(&f.Location)
	app.Flag("output", "Specifies the results to display, separated by commas").Short('o').
		Default(options.DefaultOutputOption).StringVar(&f.Output)
	app.Flag("matching-groups", "Specifies URI matching groups separated by commas").
		Short('m').PlaceHolder("PATTERN,...").StringVar(&f.MatchingGroups)
	app.Flag("filters", "Only the logs are profiled that match the conditions").
		Short('f').StringVar(&f.Filters)
}
