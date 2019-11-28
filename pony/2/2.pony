use "regex"
use "collections"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(Part2(env)) end, 512)
    env.out.print("day 2!")

  class Part2 is Processable
    let  _env : Env
    var _total : I32
    new create(env: Env) =>
      _env = env
      _total = 0
    
    fun ref process(l: String) =>
      try
        let r = Regex("(\\d+)x(\\d+)x(\\d+)")?
        let matched = r(l)?
	var l' = matched(1)?.i32()?
	var h' = matched(2)?.i32()?
	let w' = matched(3)?.i32()?
	let arr = [l';h';w']
	let sorted = Sort[Array[I32], I32](arr)

        // A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of
        // ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the
        // bow, for a total of 34 feet.  A present with dimensions 1x1x10
        // requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus
        // 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

	let wrap = (2*sorted(0)?) + (2*sorted(1)?)
	let extra = l' * h' * w'
	_total = _total + wrap + extra
      end

    fun dispose() =>
      _env.out.print(_total.string())


  class Part1 is Processable
    let  _env : Env
    var _total : I32
    new create(env: Env) =>
      _env = env
      _total = 0
    
    fun ref process(l: String) =>
      try
        let r = Regex("(\\d+)x(\\d+)x(\\d+)")?
        let matched = r(l)?
	var l' = matched(1)?.i32()?
	var h' = matched(2)?.i32()?
	let w' = matched(3)?.i32()?
	let arr = [l';h';w']
	let sorted = Sort[Array[I32], I32](arr)

	let wrap = (2*l'*w') + (2*w'*h') + (2*h'*l')
	let extra = sorted(0)? * sorted(1)?
	_total = _total + wrap + extra
      end

    fun dispose() =>
      _env.out.print(_total.string())
      
   