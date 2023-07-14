package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"seeyonerExp/core"
	"strconv"
	"sync"
)

var wg sync.WaitGroup // 并发有序的组
var mutex = &sync.Mutex{}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "漏洞检测",
	Long: `漏洞检测功能
`,
	Run: func(cmd *cobra.Command, args []string) {
		factory := new(core.IFactory)

		if vulnId == 0 {
			limiter := make(chan struct{}, 1) // 通过管道控制并发数
			arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
			for _, pro := range arr { //并发应该处于一个循环之中
				wg.Add(1)
				limiter <- struct{}{}
				go func(pro string) {
					num, _ := strconv.Atoi(pro)
					iScan := factory.NewFactory(num)
					iScan.Scan(url)
					wg.Done()
					<-limiter
				}(pro)
			}
			wg.Wait()
			os.Exit(0)
		}
		iScan := factory.NewFactory(vulnId)
		iScan.Scan(url)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&url, "targetUrl", "u", "", "targetUrl")
	scanCmd.Flags().IntVarP(&vulnId, "vulnId", "i", 0, "vulnId")
	scanCmd.MarkFlagRequired("targetUrl")
	scanCmd.MarkFlagRequired("vulnId")
}
