package main

//
//import (
//	"dbScheduleAnalysis/controllers"
//	"fmt"
//)
//
//func main() {
//	var sqlTarget =  `SELECT
//	edu_class_schedule.*,
//	GROUP_CONCAT( DISTINCT TU.id ) TA,
//	GROUP_CONCAT( DISTINCT TU.NAME ) ta_name
//FROM
//	(
//	SELECT
//		edu_class_schedule.id,
//		edu_class_schedule.class_id,
//		edu_class_schedule.classroom_id,
//		edu_class_schedule.app_start_time,
//		edu_class_schedule.app_end_time,
//		edu_class_schedule.study_state,
//		unit.id unit_new_version_id,
//		CONCAT( "", unit.NAME ) unit,
//		lession.id lesson_new_version_id,
//		CONCAT( "", lession.NAME ) lession,
//		schedule_textbook.unit_id,
//		schedule_textbook.tb_id,
//		edu_class_schedule.remark,
//		base_otm_classroom.NAME,
//		base_class_manage.class_number,
//		base_class_manage.cm_state,
//		base_class_manage.class_type,
//		prod_text_book.NAME tb_name,
//		edu_class_schedule.tea_id FT,
//		FU.NAME ft_name
//	FROM
//		edu_class_schedule
//		JOIN edu_course_schedule_level unit ON unit.id =(
//		CASE
//
//				WHEN edu_class_schedule.unit < 20 THEN
//				- edu_class_schedule.unit ELSE edu_class_schedule.unit
//			END
//			)
//			JOIN edu_course_schedule_level lession ON lession.id =(
//			CASE
//
//					WHEN edu_class_schedule.lession < 20 THEN
//					- 100-edu_class_schedule.lession ELSE edu_class_schedule.lession
//				END
//				)
//				LEFT JOIN (
//				SELECT
//					edu_class_schedule.id sc_id,
//					stu_class_type_textbook.tb_id,
//					(
//					SELECT
//						id
//					FROM
//						tch_unit
//					WHERE
//						unit_name = (
//						CASE
//
//								WHEN edu_class_schedule.unit < 20 THEN
//								CONCAT( "unit", edu_class_schedule.unit ) ELSE ( SELECT NAME FROM edu_course_schedule_level WHERE id = edu_class_schedule.unit LIMIT 1 )
//							END
//							)) unit_id
//					FROM
//						edu_class_schedule
//						JOIN stu_class ON stu_class.cm_id = edu_class_schedule.class_id
//						JOIN stu_class_type_textbook ON stu_class_type_textbook.order_id = stu_class.of_id
//					WHERE
//						edu_class_schedule.id = 182070
//					GROUP BY
//						edu_class_schedule.id
//					) schedule_textbook ON schedule_textbook.sc_id = edu_class_schedule.id
//					LEFT JOIN base_otm_classroom ON edu_class_schedule.classroom_id = base_otm_classroom.id
//					LEFT JOIN base_class_manage ON edu_class_schedule.class_id = base_class_manage.id
//					LEFT JOIN prod_text_book ON base_class_manage.tb_id = prod_text_book.id
//					LEFT JOIN sys_user FU ON FU.id = edu_class_schedule.tea_id
//				WHERE
//					edu_class_schedule.id = 182070
//					LIMIT 1
//				) edu_class_schedule
//				LEFT JOIN edu_class_allot_customer ecac ON edu_class_schedule.class_id = ecac.class_id
//				AND ecac.allot_state = 17
//				AND ecac.del_flag = 0
//				LEFT JOIN sys_user TU ON TU.id = ecac.customer_id
//		GROUP BY
//	edu_class_schedule.id`
//	var result = controllers.ParseSqlSentenceAboutOn(sqlTarget)
//	fmt.Printf("result: %v\n",result )
//}
