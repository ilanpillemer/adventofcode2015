
actor Main
  new create(env: Env) =>
    var count: I32  = 0
    var pos: I32 = 0
    var done: Bool = false
    env.input(
      object iso is InputNotify

        fun ref apply(data: Array[U8] iso) =>
	    if done then return end
            var data':Array[U8] ref = consume data
            for i in data'.values() do
	     pos = pos + 1
             match i
             | ')' => count = count - 1
             | '(' => count = count + 1
             end
	     if count < 0 then
	      env.out.print("basement: " + pos.string())
	      env.input.dispose()
	      done = true
	      return
	     end
            end

        fun ref dispose() =>
          env.out.print(count.string())
      end,
      512)

