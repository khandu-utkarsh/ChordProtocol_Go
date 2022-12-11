package chord

import (
	"bytes"
	"crypto/sha1"
	"math/big"
)

type HashId struct {
	id []byte
}

// !Auxilarry function
func Generate_Hash(inp []byte) HashId {
	h := sha1.New()
	h.Write(inp)
	bs := h.Sum(nil)

	hash := HashId{id: bs}
	return hash
}

//!Change these before submitting
func GetCurrentProcessIPAddress() string {
	return "AurojitPanda"
}

func GetCurrentProcessPort() string {
	return ":Sucks"
}



// Implement the comparison function
func IsIdBetweenRange_RightEnd_Inclusive(key HashId, min HashId, max HashId) bool {
	key_hash := key.id
	min_hash := min.id
	max_hash := max.id
	firstCond := bytes.Compare(min_hash, key_hash) < 0    // -> True if min < key
	seocondCond := bytes.Compare(key_hash, max_hash) <= 0 // -> True if key <= max

	if firstCond && seocondCond {
		return true
	}
	return false
}

func IsIdBetweenRange_RightEnd_Exclusive(key HashId, min HashId, max HashId) bool {
	key_hash := key.id
	min_hash := min.id
	max_hash := max.id

	firstCond := bytes.Compare(min_hash, key_hash) < 0   // -> True if min < key
	seocondCond := bytes.Compare(key_hash, max_hash) < 0 // -> True if key < max

	if firstCond && seocondCond {
		return true
	}
	return false
}


func GenerateHashIdForFingerIndex(n HashId, indexOfPower int) (HashId) {
	var ret_hash_id HashId

	//!Generate a big interger for n
	var n_big_int big.Int
	n_big_int.SetBytes(n.id)
	
	//!Generate a big interget for 2 to power indexOfPower
	var power_base_int big.Int
	power_base_int.SetUint64(2)

	var power_exponent_int big.Int
	power_exponent_int.SetUint64(uint64(indexOfPower))

	var power_out big.Int
	power_out.Exp(&power_base_int, &power_exponent_int, nil)

	var sum_out big.Int
	sum_out.Add(&n_big_int, &power_out)

	var power_exponent_m_int big.Int
	power_exponent_m_int.SetUint64(uint64(m))

	var two_power_m_int big.Int
	two_power_m_int.Exp(&power_base_int, &power_exponent_m_int, nil)

	var final_out big.Int
	final_out.Mod(&sum_out, &two_power_m_int)


	output :=make([]byte, m, m)
	out := final_out.FillBytes(output)

	ret_hash_id.id = out
	return ret_hash_id
}