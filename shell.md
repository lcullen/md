cat /data/service_logs/services/call.sale.log_2019-03-18 | grep sc_sql_log | awk -F , '{print $3}' | sed "s/@\[php@\]//g" >> sc_sql_raw.sql


cat call.sale.log_2019-02-25 | grep sc_sql_log | awk -F , '{print $3}' | sed "s/@\[php@\]//g" | sed "s/'[^']*'//g"  |  sed "s/[0-9]//g" >> /home/rd/sc_sql_log