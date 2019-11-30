use "crypto"

actor Main
  var count: U64 = 0
  let prefix: String = "yzbqklnj"
  let _env: Env
  new create(env: Env) =>
    _env = env
    _env.out.print("day4 2015 in Pony")
    _env.out.print("part 2")
    search()

  //Cant do a tight loop as Pony wont gc during a behaviour
  //so it will run out of heap! 
  be search() =>
    var cand = ToHexString(MD5(prefix + count.string()))
    if cand.substring(0,6) == "000000" then
      _env.out.print(count.string())
      _env.out.print(cand)
    else
      count = count + 1
      search()
    end
