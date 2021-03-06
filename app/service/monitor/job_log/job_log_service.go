// ==========================================================================
// 云捷GO自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/service/module/log/log_service.go
// 生成人：yunjie
// ==========================================================================
package job_log

import (
	logModel "yj-app/app/model/monitor/job_log"
	"yj-app/app/service/utils/convert"
	"yj-app/app/service/utils/excel"
	"yj-app/app/service/utils/page"
)

//根据主键查询数据
func SelectRecordById(id int64) (*logModel.Entity, error) {
	return logModel.FindOne("job_log_id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := logModel.Delete("job_log_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := logModel.Delete("job_log_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//根据条件查询数据
func SelectListAll(params *logModel.SelectPageReq) (*[]logModel.Entity, error) {
	return logModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListByPage(params *logModel.SelectPageReq) (*[]logModel.Entity, *page.Paging, error) {
	return logModel.SelectListByPage(params)
}

// 导出excel
func Export(param *logModel.SelectPageReq) (string, error) {
	result, err := logModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{"任务日志ID", "任务名称", "任务组名", "调用目标字符串", "日志信息", "执行状态（0正常 1失败）", "异常信息", "创建时间"}

	url, err := excel.DownlaodExcel(head, result)

	if err != nil {
		return "", err
	}

	return url, nil
}
