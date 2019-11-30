use "debug"
use "collections/persistent"
use "promises"

actor Main
  new create(env: Env) =>
    env.out.print("day3")
    env.input(Director,512)

class iso Director is InputNotify
 let human: Deliverer = Deliverer
 let robot: Deliverer = Deliverer
 var turn:Bool = false

 fun ref apply(data': Array[U8] iso) =>
   var data: Array[U8] ref = consume data'
   for d in data.values() do
       if turn then
         robot.move(d)
       else
         human.move(d)
       end
       turn = not turn
   end

  fun dispose() =>
    let p1 = Promise[Set[String]] 
    let p2 = Promise[Set[String]]
    Promises[Set[String] val].join([p1;p2].values())
      .next[None]({(a: Array[Set[String]] val) =>
        var dedup = Set[String]
        for m in a.values() do
          for k in m.values() do
            dedup = dedup.add(k)
          end
        end
        Debug.out(dedup.size().string())
       })
    robot.combine(p1)
    human.combine(p2)


actor Deliverer
  var x: I32 = 0
  var y: I32 = 0
  
  var map: Set[String] =  Set[String] 

  new create() =>
    map = map.add("0,0")

  be move(d: U8) =>
    match d
    | '<' => x = x - 1
    | '>' => x = x + 1
    | '^' => y = y - 1
    | 'v' => y = y + 1
    end
    var pos = x.string() + "," + y.string()
    map = map.add(pos)

  be combine(p: Promise[Set[String]]) =>
    p(map)

  
