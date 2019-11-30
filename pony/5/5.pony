use "debug"
use "regex"

actor Main
  new create(env: Env) =>
    env.out.print("pony day 5 2015")
    env.input(recover LineNotify(A) end,512)

class A
  var total:U32 = 0
  fun ref apply(s: String) =>
    var isnice = false
//    if rule3(s) and rule1(s) and rule2(s) then
    if rule4(s) and rule5(s) then    
      isnice = true
      total = total + 1
    end

  fun dispose() =>
    Debug.out(total.string())

// Part 1
// It contains at least three vowels (aeiou only), like aei, xazegov, or
// aeiouaeiouaeiou.  It contains at least one letter that appears twice
// in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).  It
// does not contain the strings ab, cd, pq, or xy, even if they are part
// of one of the other requirements


  fun rule1(s: String): Bool =>
    var count: U32 = 0
    for i in s.values() do
    match i
      | 'a' => count = count + 1
      | 'e' => count = count + 1
      | 'i' => count = count + 1
      | 'o' => count = count + 1
      | 'u' => count = count + 1
      end
    end
    count >= 3

  fun rule3(s: String): Bool =>
    if s.contains("ab") then return false end
    if s.contains("cd") then return false end
    if s.contains("pq") then return false end
    if s.contains("xy") then return false end
    true

  fun rule2(s: String): Bool =>
      try
       let r = Regex("(.)\\1")?
       r(s)?
       true
      else
        false
      end

//  Part 2
//  It contains a pair of any two letters that appears at least twice
//  in the string without overlapping, like xyxy (xy) or aabcdefgaa
//  (aa), but not like aaa (aa, but it overlaps).  It contains at
//  least one letter which repeats with exactly one letter between
//  them, like xyx, abcdefeghi (efe), or even aaa

  fun rule4(s: String): Bool =>
    try
    let r = Regex("(..).*\\1")?
    r(s)?
    true
    else
      false
    end

  fun rule5(s: String): Bool =>
    try
    let r = Regex("(.).\\1")?
    r(s)?
    true
    else
      false
    end
  