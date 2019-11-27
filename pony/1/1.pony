
actor Main
  new create(env: Env) =>
    var count: I32  = 0
    env.input(
      object iso is InputNotify

        fun ref apply(data: Array[U8] iso) =>
            var data':Array[U8] ref = consume data
            for i in data'.values() do
             match i
             | ')' => count = count - 1
             | '(' => count = count + 1
             end
            end

        fun ref dispose() =>
          env.out.print(count.string())
      end,
      512)

