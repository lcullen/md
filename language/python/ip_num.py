#! /usr/bin/env python
# -*- coding: utf-8 -*-
class Solution:
    def _isIpNum(self,num):
        if 0<=int(num)<=255 and str(int(num)) == num:
            return True
        return False

    def _restoreIpAddress(self, leftCnt, index, s, nowIp, result):
        if leftCnt == 0: #边界检查如果已经没有剩余的"."的数量 则返回
            if index == len(s): # 并且 刚好到最后一个下标 表示满足要求
                result.append(nowIp)
                return

        for i in range(index+1, len(s)+1):
            if self._isIpNum(s[index:i]):
                nowIpTmp = s[index:i] if nowIp == "" else nowIp +"." + s[index:i]
                self._restoreIpAddress(leftCnt-1,i,s, nowIpTmp, result)
            else:
                break

    def restoreIpAddress(self, s):
        result = list()
        if not s and 4 > len(s) > 12:
            return result

        self._restoreIpAddress(4,0,s,"",result)
        return result


t = Solution()
print t.restoreIpAddress("25525511135")
