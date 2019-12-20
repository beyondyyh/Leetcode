package linkedList

import "testing"

func build_isPalindrome_case() (cases []*ListNode) { // {{{
	head1 := new(ListNode)
	head1.Val = 1
	c1l2 := new(ListNode)
	c1l2.Val = 2
	c1l3 := new(ListNode)
	c1l3.Val = 3
	c1l4 := new(ListNode)
	c1l4.Val = 2
	c1l5 := new(ListNode)
	c1l5.Val = 1
	head1.Next = c1l2
	c1l2.Next = c1l3
	c1l3.Next = c1l4
	c1l4.Next = c1l5

	head2 := new(ListNode)
	head2.Val = 1
	c2l2 := new(ListNode)
	c2l2.Val = 2
	c2l3 := new(ListNode)
	c2l3.Val = 2
	c2l4 := new(ListNode)
	c2l4.Val = 1
	head2.Next = c2l2
	c2l2.Next = c2l3
	c2l3.Next = c2l4

	head3 := new(ListNode)
	head3.Val = 1
	c3l2 := new(ListNode)
	c3l2.Val = 2
	c3l3 := new(ListNode)
	c3l3.Val = 3
	c3l4 := new(ListNode)
	c3l4.Val = 2
	head3.Next = c3l2
	c3l2.Next = c3l3
	c3l3.Next = c3l4

	head4 := new(ListNode)
	head4.Val = 1

	cases = append(cases, head1, head2, head3, head4)
	return
} // }}}

func Test_isPalindrome(t *testing.T) { // {{{
	cases := build_isPalindrome_case()
	tests := []struct {
		name   string
		intput *ListNode
		want   bool
	}{
		{
			name:   "基数回文",
			intput: cases[0],
			want:   true,
		},
		{
			name:   "偶数回文",
			intput: cases[1],
			want:   true,
		},
		{
			name:   "非回文",
			intput: cases[2],
			want:   false,
		},
		{
			name:   "单节点回文",
			intput: cases[3],
			want:   true,
		},
		{
			name:   "空节点回文",
			intput: nil,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.intput); got != tt.want {
				t.Errorf("isPalindrome() = %v, want: %v", got, tt.want)
			}
		})
	}
} // }}}
