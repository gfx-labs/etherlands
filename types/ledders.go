package types

func Create24Name(name string) [24]byte {
  if name == "" {
    return [24]byte{}
  }
  fixed_name := [24]byte{}
  for i := 0; (i < len(name)) && (i < 24); i++ {
    if(i >= len(name)){
      fixed_name[i] = 0x00;
    }else{
      fixed_name[i] = LedderString(string(name[i]));
    }
  }
  return fixed_name;
}

func Parse24Name(nickname [24]byte) string{
  base := ""
  for _, b := range nickname {
    if b == 0x00 {return base}
    if b >= 0x40 {return base}
    base = base + ParseLedder(b);
  }
  return base;
}

func ParseLedder(b byte) string {
  if val, ok := keycode[b]; ok {
    return val;
  }
  return ""
}
func LedderString(s string) byte {
  if val, ok := codekey[s]; ok {
    return val;
  }
  return 0x00;
}


var keycode = map[byte]string{
  0x00 : "",
  0x01 : "1",
  0x02 : "2",
  0x03 : "3",
  0x04 : "4",
  0x05 : "5",
  0x06 : "6",
  0x07 : "7",
  0x08 : "8",
  0x09 : "9",
  0x0a : "0",
  0x0b : "a",
  0x0c : "b",
  0x0d : "c",
  0x0e : "d",
  0x0f : "e",
  0x10 : "f",
  0x11 : "g",
  0x12 : "h",
  0x13 : "i",
  0x14 : "j",
  0x15 : "k",
  0x16 : "l",
  0x17 : "m",
  0x18 : "n",
  0x19 : "o",
  0x1a : "p",
  0x1b : "q",
  0x1c : "r",
  0x1d : "s",
  0x1e : "t",
  0x1f : "u",
  0x20 : "v",
  0x21 : "w",
  0x22 : "x",
  0x23 : "y",
  0x24 : "z",
  0x25 : "A",
  0x26 : "B",
  0x27 : "C",
  0x28 : "D",
  0x29 : "E",
  0x2a : "F",
  0x2b : "G",
  0x2c : "H",
  0x2d : "I",
  0x2e : "J",
  0x2f : "K",
  0x30 : "L",
  0x31 : "M",
  0x32 : "N",
  0x33 : "O",
  0x34 : "P",
  0x35 : "Q",
  0x36 : "R",
  0x37 : "S",
  0x38 : "T",
  0x39 : "U",
  0x3a : "V",
  0x3b : "W",
  0x3c : "X",
  0x3d : "Y",
  0x3e : "Z",
  0x3f : "_",
  0x40 : "",
  0x41 : "",
}


var codekey = map[string]byte{
  "" : 0x00,
  "1": 0x01,
  "2": 0x02,
  "3": 0x03,
  "4": 0x04,
  "5": 0x05,
  "6": 0x06,
  "7": 0x07,
  "8": 0x08,
  "9": 0x09,
  "0": 0x0a,
  "a": 0x0b,
  "b": 0x0c,
  "c": 0x0d,
  "d": 0x0e,
  "e": 0x0f,
  "f": 0x10,
  "g": 0x11,
  "h": 0x12,
  "i": 0x13,
  "j": 0x14,
  "k": 0x15,
  "l": 0x16,
  "m": 0x17,
  "n": 0x18,
  "o": 0x19,
  "p": 0x1a,
  "q": 0x1b,
  "r": 0x1c,
  "s": 0x1d,
  "t": 0x1e,
  "u": 0x1f,
  "v": 0x20,
  "w": 0x21,
  "x": 0x22,
  "y": 0x23,
  "z": 0x24,
  "A": 0x25,
  "B": 0x26,
  "C": 0x27,
  "D": 0x28,
  "E": 0x29,
  "F": 0x2a,
  "G": 0x2b,
  "H": 0x2c,
  "I": 0x2d,
  "J": 0x2e,
  "K": 0x2f,
  "L": 0x30,
  "M": 0x31,
  "N": 0x32,
  "O": 0x33,
  "P": 0x34,
  "Q": 0x35,
  "R": 0x36,
  "S": 0x37,
  "T": 0x38,
  "U": 0x39,
  "V": 0x3a,
  "W": 0x3b,
  "X": 0x3c,
  "Y": 0x3d,
  "Z": 0x3e,
  "_": 0x3f,
}