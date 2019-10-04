1. 边界检查
2. 最后写代码
```php
 function isCircle($node) {
        $slowNode = $quickNode = $node;
        while($quickNode  != null) {
            $slowNode = $slowNode->next;
            if($quickNode->next == null) return false;
            $quickNode = $quickNode->next->next;
            if($slowNode->equals($quickNode)) return true;
        }
    }

    function findNumInDoubleArr($arr, $num) {
        if(count($arr) == 0) return false;
        $positionX = count($arr[0]) - 1;
        $positionY = 0;
        while($positionX >= 0 && $positionY < count($arr)) {
            $positionNum = $arr[$positionX][$positionY];
            if($positionNum == $num) return true;
            if($positionNum < $num) $positionY ++;
            else $positionX --;
        }
        return false;
    }

    //查找单调不递减 旋转数组最小值
    function findGap($arr) {
        if(!is_array($arr) || count($arr) == 0 ) return 0;
        $begin = 0;
        $end = count($arr) - 1;
        while ($begin <= $end) {
            if($end - $begin == 1) break;
            $middle = ($begin + $end) / 2;
            $middleV = $arr[$middle];
            if($middleV > $arr[$begin]) $begin = $middle;
            elseif($middle < $arr[$begin]) $end = $middle;
        }
        return $arr[$end];
    }

    function frogJump($n) {
        $dp[0] = 0;
        $dp[1] = 1;
        $dp[2] = 2;
        $i = 2;
        if($n <= 2) return $dp[$n];
        while($n != $i) {
            $dp[$i] = $dp[$i-1] + $dp[$i-2];
        }
        return $dp[$n];
    }

    function splitOdd($arr) {
        if(!is_array($arr) || count($arr) <= 1) return $arr;
        $idxV = $arr[0];
        $begin = 0;
        $end = count($arr) - 1;
        while($begin <= $end) {
            while($arr[$end] % 2 == 0 && $end > $begin) $end --;
            while($arr[$begin] % 2 == 1 && $end < $begin) $begin ++;
            swap($arr, $begin, $end);
        }
    }

    function mergeLinkList($node1, $node2) {
        if($node1 == null) return $node2;
        if($node2 == null) return $node1;
        $newHead = $tmp = new Node();
        while($node1 != null && $node2 != null) {
            if ($node1->value > $node2->value){
                $tmp->next = $node2;
                $node2 = $node2->next;
            } else {
                $tmp->next = $node1;
                $node1 = $node1->next;
            }
            $tmp = $tmp->next;
        }

        if($node2 != null) {
            $newHead->next = $node2;
        }

        if($node1 != null) {
            $newHead->next = $node1;
        }
        return $newHead->next;
    }

    function isIncludeTree($bigTree, $small) {
        if($bigTree == null) return false;
        if($small == null) return true;

        if($bigTree->value == $small->value) {
            return $this->isIncludeTree($bigTree->left, $small->left)
                && $this->isIncludeTree($bigTree->right, $small->right);
        } else {
            return $this->isIncludeTree($bigTree->left, $small)
            || $this->isIncludeTree($bigTree->right, $small);
        }
    }

    //sumInTreeRoute 树中某一路劲 sum
    function sumInTreeRoute($treeNode, $sum) {
        if($treeNode == null) return null;
        $s = new Stack();
        if($sum == 0 && $treeNode->value == 0) {

        }
    }

    //先序遍历 + stack
    function sumInTreeRouterInner($stack, $treeNode, $sum, $cursum) {
        //if($treeNode == null) return $stack;
        $stack->push($treeNode);
        $cursum += $treeNode->val;
        $isLeaf = $treeNode->left == null && $treeNode->right == null;
        if($isLeaf && $cursum == $sum) {
            print_r($stack);
        }

        if($treeNode->left) {
            $this->sumInTreeRouterInner($stack, $treeNode->left, $sum, $cursum);
        }
        if($treeNode->right) {
            $this->sumInTreeRouterInner($stack, $treeNode->right, $sum, $cursum);
        }

        //每次函数执行到这里 函数栈 弹出; 回溯 原值
        $cursum -= $treeNode->val;
        $stack->pop();
    }

    //中序遍历 + Linked list 将一颗二叉搜索树 转化成有序 链表
    function revertTreeToLink($treeNode, $curLinkedNode) {

        if($treeNode->left != null) {
            $curLinkedNode->before = $this->revertTreeToLink($treeNode->left, $curLinkedNode);
        }
        if($curLinkedNode != null) {
            $curLinkedNode->after = $treeNode;
        }
        if($treeNode->right != null) {
            $curLinkedNode->after = $this->revertTreeToLink($treeNode->right, $curLinkedNode);
        }

        return $treeNode;
    }

```