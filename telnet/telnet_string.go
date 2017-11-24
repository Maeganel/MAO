// Code generated by "stringer -type tnSeq -output telnet_string.go"; DO NOT EDIT.

package telnet

import "fmt"

const (
	_tnSeq_name_0 = "NULECHO"
	_tnSeq_name_1 = "SGA"
	_tnSeq_name_2 = "STTMBELBSHTLF"
	_tnSeq_name_3 = "FFCR"
	_tnSeq_name_4 = "TTEOR"
	_tnSeq_name_5 = "WSTSRFCLM"
	_tnSeq_name_6 = "EV"
	_tnSeq_name_7 = "CMP1CMP2"
	_tnSeq_name_8 = "ATCPGMCP"
	_tnSeq_name_9 = "SENOPDMBRKIPAOAYTECELGASBWILLWONTDODONTIAC"
)

var (
	_tnSeq_index_0 = [...]uint8{0, 3, 7}
	_tnSeq_index_1 = [...]uint8{0, 3}
	_tnSeq_index_2 = [...]uint8{0, 2, 4, 7, 9, 11, 13}
	_tnSeq_index_3 = [...]uint8{0, 2, 4}
	_tnSeq_index_4 = [...]uint8{0, 2, 5}
	_tnSeq_index_5 = [...]uint8{0, 2, 4, 7, 9}
	_tnSeq_index_6 = [...]uint8{0, 2}
	_tnSeq_index_7 = [...]uint8{0, 4, 8}
	_tnSeq_index_8 = [...]uint8{0, 4, 8}
	_tnSeq_index_9 = [...]uint8{0, 2, 5, 7, 10, 12, 14, 17, 19, 21, 23, 25, 29, 33, 35, 39, 42}
)

func (i tnSeq) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _tnSeq_name_0[_tnSeq_index_0[i]:_tnSeq_index_0[i+1]]
	case i == 3:
		return _tnSeq_name_1
	case 5 <= i && i <= 10:
		i -= 5
		return _tnSeq_name_2[_tnSeq_index_2[i]:_tnSeq_index_2[i+1]]
	case 12 <= i && i <= 13:
		i -= 12
		return _tnSeq_name_3[_tnSeq_index_3[i]:_tnSeq_index_3[i+1]]
	case 24 <= i && i <= 25:
		i -= 24
		return _tnSeq_name_4[_tnSeq_index_4[i]:_tnSeq_index_4[i+1]]
	case 31 <= i && i <= 34:
		i -= 31
		return _tnSeq_name_5[_tnSeq_index_5[i]:_tnSeq_index_5[i+1]]
	case i == 36:
		return _tnSeq_name_6
	case 85 <= i && i <= 86:
		i -= 85
		return _tnSeq_name_7[_tnSeq_index_7[i]:_tnSeq_index_7[i+1]]
	case 200 <= i && i <= 201:
		i -= 200
		return _tnSeq_name_8[_tnSeq_index_8[i]:_tnSeq_index_8[i+1]]
	case 240 <= i && i <= 255:
		i -= 240
		return _tnSeq_name_9[_tnSeq_index_9[i]:_tnSeq_index_9[i+1]]
	default:
		return fmt.Sprintf("tnSeq(%d)", i)
	}
}