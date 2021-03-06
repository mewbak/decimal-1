
//line parser.rl:1
package suite

import (
    "fmt"
    "strconv"
)

func parseCase(data []byte) (c Case, err error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var (
        ok   bool // for mode and op
        mark int
        exc  Exception
    )

    
//line parser.go:21
const parser_start int = 1
const parser_first_final int = 50
const parser_error int = 0

const parser_en_main int = 1


//line parser.go:29
	{
	cs = parser_start
	}

//line parser.go:34
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 50:
		goto st_case_50
	case 14:
		goto st_case_14
	case 51:
		goto st_case_51
	case 15:
		goto st_case_15
	case 52:
		goto st_case_52
	case 16:
		goto st_case_16
	case 53:
		goto st_case_53
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 54:
		goto st_case_54
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 98:
			goto tr0
		case 100:
			goto tr0
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line parser.rl:20
 mark = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:173
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
tr2:
//line parser.rl:21
 c.Prefix = string(data[mark:p]) 
//line parser.rl:20
 mark = p 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:189
		switch data[p] {
		case 37:
			goto tr3
		case 42:
			goto tr4
		case 43:
			goto tr3
		case 45:
			goto tr3
		case 47:
			goto tr3
		case 61:
			goto tr7
		case 63:
			goto tr8
		case 76:
			goto tr3
		case 78:
			goto tr9
		case 83:
			goto tr3
		case 86:
			goto tr3
		case 99:
			goto tr10
		case 101:
			goto tr11
		case 113:
			goto tr12
		case 114:
			goto tr13
		case 115:
			goto tr14
		case 126:
			goto tr3
		}
		switch {
		case data[p] < 60:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 62:
			if 64 <= data[p] && data[p] <= 65 {
				goto tr3
			}
		default:
			goto tr6
		}
		goto st0
tr3:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:255
		if data[p] == 32 {
			goto tr15
		}
		goto st0
tr15:
//line parser.rl:28

            c.Op, ok = valToOp[string(data[mark:p])]
            if !ok {
                return c, fmt.Errorf("invalid op: %q", data[mark:p])
            }
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser.go:274
		switch data[p] {
		case 48:
			goto tr16
		case 61:
			goto tr17
		}
		if 60 <= data[p] && data[p] <= 62 {
			goto tr16
		}
		goto st0
tr16:
//line parser.rl:20
 mark = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:294
		if data[p] == 32 {
			goto tr18
		}
		goto st0
tr18:
//line parser.rl:34

	    c.Mode, ok = valToMode[string(data[mark:p])]
	    if !ok {
		return c, fmt.Errorf("invalid mode: %q", data[mark:p])
	    }
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:313
		switch data[p] {
		case 43:
			goto tr19
		case 45:
			goto tr20
		case 73:
			goto tr22
		case 81:
			goto tr23
		case 83:
			goto tr23
		case 105:
			goto tr24
		case 111:
			goto tr25
		case 122:
			goto tr25
		}
		switch {
		case data[p] > 57:
			if 117 <= data[p] && data[p] <= 120 {
				goto tr25
			}
		case data[p] >= 48:
			goto tr21
		}
		goto st0
tr19:
//line parser.rl:20
 mark = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:350
		switch data[p] {
		case 73:
			goto st21
		case 105:
			goto st21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
tr21:
//line parser.rl:20
 mark = p 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:370
		switch data[p] {
		case 32:
			goto tr28
		case 46:
			goto st24
		case 101:
			goto st26
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
tr28:
//line parser.rl:47
 c.Inputs = append(c.Inputs, Data(data[mark:p])) 
	goto st10
tr50:
//line parser.rl:40

	    exc, ok = valToException[string(data[mark:p])]
	    if !ok {
                return c, fmt.Errorf("invalid trap exception: %q", data[mark:p])
	    }
            c.Trap |= exc
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:402
		switch data[p] {
		case 43:
			goto tr19
		case 45:
			goto tr20
		case 73:
			goto tr22
		case 81:
			goto tr23
		case 83:
			goto tr23
		case 105:
			goto tr22
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr21
		}
		goto st0
tr20:
//line parser.rl:20
 mark = p 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parser.go:430
		switch data[p] {
		case 62:
			goto st12
		case 73:
			goto st21
		case 105:
			goto st21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 32 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 35:
			goto tr33
		case 43:
			goto tr34
		case 45:
			goto tr34
		case 73:
			goto tr36
		case 81:
			goto tr33
		case 83:
			goto tr33
		case 105:
			goto tr36
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr35
		}
		goto st0
tr33:
//line parser.rl:20
 mark = p 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line parser.go:486
		if data[p] == 32 {
			goto tr62
		}
		goto st0
tr62:
//line parser.rl:48
 c.Output = Data(data[mark:p]) 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:500
		switch data[p] {
		case 105:
			goto tr37
		case 111:
			goto tr37
		case 122:
			goto tr37
		}
		if 117 <= data[p] && data[p] <= 120 {
			goto tr37
		}
		goto st0
tr37:
//line parser.rl:20
 mark = p 
	goto st51
tr63:
//line parser.rl:49

	    exc, ok = valToException[string(data[mark:p])]
	    if !ok {
		return c, fmt.Errorf("invalid result exception: %q", data[mark:p])
	    }
            c.Excep |= exc
        
//line parser.rl:20
 mark = p 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line parser.go:534
		switch data[p] {
		case 105:
			goto tr63
		case 111:
			goto tr63
		case 122:
			goto tr63
		}
		if 117 <= data[p] && data[p] <= 120 {
			goto tr63
		}
		goto st0
tr34:
//line parser.rl:20
 mark = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:556
		switch data[p] {
		case 73:
			goto st19
		case 105:
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st52
		}
		goto st0
tr35:
//line parser.rl:20
 mark = p 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line parser.go:576
		switch data[p] {
		case 32:
			goto tr62
		case 46:
			goto st16
		case 101:
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st52
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 32:
			goto tr62
		case 101:
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 45 {
			goto st18
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st54
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= data[p] && data[p] <= 57 {
			goto st54
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 32 {
			goto tr62
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st54
		}
		goto st0
tr36:
//line parser.rl:20
 mark = p 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:655
		switch data[p] {
		case 78:
			goto st20
		case 110:
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 70:
			goto st50
		case 102:
			goto st50
		}
		goto st0
tr22:
//line parser.rl:20
 mark = p 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:684
		switch data[p] {
		case 78:
			goto st22
		case 110:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 70:
			goto st23
		case 102:
			goto st23
		}
		goto st0
tr23:
//line parser.rl:20
 mark = p 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:713
		if data[p] == 32 {
			goto tr28
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 32:
			goto tr28
		case 101:
			goto st26
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 45 {
			goto st27
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st28
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if 48 <= data[p] && data[p] <= 57 {
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 32 {
			goto tr28
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st28
		}
		goto st0
tr24:
//line parser.rl:20
 mark = p 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:784
		switch data[p] {
		case 32:
			goto tr50
		case 78:
			goto st22
		case 105:
			goto tr51
		case 110:
			goto st22
		case 111:
			goto tr51
		case 122:
			goto tr51
		}
		if 117 <= data[p] && data[p] <= 120 {
			goto tr51
		}
		goto st0
tr25:
//line parser.rl:20
 mark = p 
	goto st30
tr51:
//line parser.rl:40

	    exc, ok = valToException[string(data[mark:p])]
	    if !ok {
                return c, fmt.Errorf("invalid trap exception: %q", data[mark:p])
	    }
            c.Trap |= exc
        
//line parser.rl:20
 mark = p 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:824
		switch data[p] {
		case 32:
			goto tr50
		case 105:
			goto tr51
		case 111:
			goto tr51
		case 122:
			goto tr51
		}
		if 117 <= data[p] && data[p] <= 120 {
			goto tr51
		}
		goto st0
tr17:
//line parser.rl:20
 mark = p 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:848
		switch data[p] {
		case 48:
			goto st6
		case 94:
			goto st6
		}
		goto st0
tr4:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:872
		switch data[p] {
		case 32:
			goto tr15
		case 45:
			goto st4
		}
		goto st0
tr6:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line parser.go:896
		switch data[p] {
		case 65:
			goto st4
		case 67:
			goto st4
		}
		goto st0
tr7:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:920
		if data[p] == 113 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 117 {
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 97 {
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 110 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 116 {
			goto st4
		}
		goto st0
tr8:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line parser.go:977
		switch data[p] {
		case 32:
			goto tr15
		case 45:
			goto st4
		case 48:
			goto st4
		case 78:
			goto st4
		case 102:
			goto st4
		case 105:
			goto st4
		case 110:
			goto st4
		case 115:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 32:
			goto tr15
		case 78:
			goto st4
		}
		goto st0
tr9:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line parser.go:1025
		switch data[p] {
		case 97:
			goto st4
		case 100:
			goto st4
		case 117:
			goto st4
		}
		goto st0
tr10:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line parser.go:1051
		switch data[p] {
		case 100:
			goto st43
		case 102:
			goto st44
		case 105:
			goto st43
		case 112:
			goto st4
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 102 {
			goto st4
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 100:
			goto st4
		case 102:
			goto st4
		case 105:
			goto st4
		}
		goto st0
tr11:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line parser.go:1102
		if data[p] == 113 {
			goto st4
		}
		goto st0
tr12:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line parser.go:1123
		switch data[p] {
		case 67:
			goto st4
		case 117:
			goto st36
		}
		goto st0
tr13:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line parser.go:1147
		if data[p] == 102 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 105 {
			goto st4
		}
		goto st0
tr14:
//line parser.rl:22

            c.Prec, err = strconv.Atoi(string(data[mark:p]))
            if err != nil {
                return c, err
            }
        
//line parser.rl:20
 mark = p 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line parser.go:1177
		if data[p] == 67 {
			goto st4
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 50, 52, 53, 54:
//line parser.rl:48
 c.Output = Data(data[mark:p]) 
		case 51:
//line parser.rl:49

	    exc, ok = valToException[string(data[mark:p])]
	    if !ok {
		return c, fmt.Errorf("invalid result exception: %q", data[mark:p])
	    }
            c.Excep |= exc
        
//line parser.go:1252
		}
	}

	_out: {}
	}

//line parser.rl:141

    return c, nil
}
