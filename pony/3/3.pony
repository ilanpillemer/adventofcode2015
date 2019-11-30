use "debug"
use "collections/persistent"
use "promises"

actor Main
  new create(env: Env) =>
    env.out.print("day3")
    env.input(Director,512)

class iso Director is InputNotify
 let santa:Santa = Santa
 let robo:Santa = Santa
 var turn:Bool = false

 fun ref apply(data': Array[U8] iso) =>
   var data: Array[U8] ref = consume data'
   for d in data.values() do
       if turn then
         santa.move(d)
       else
         robo.move(d)
       end
       turn = not turn
   end

  fun dispose() =>
    let p1 = Promise[Map[String,U32]] 
    let p2 = Promise[Map[String,U32]]
    Promises[Map[String,U32] val].join([p1;p2].values())
      .next[None]({(a: Array[Map[String,U32]] val) =>
	var dedup = Map[String,U32]
	for m in a.values() do
	  for k in m.keys() do
	    dedup = dedup.update(k,0)
          end
	end
	Debug.out(dedup.size().string())
       })
    robo.combine(p1)
    santa.combine(p2)


actor Santa
  var x: I32 = 0
  var y: I32 = 0
  
  var map: Map[String,U32] =  Map[String,U32] 

  new create() =>
    map = map.update("0,0",1)

  be move(d: U8) =>
    match d
    | '<' => x = x - 1
    | '>' => x = x + 1
    | '^' => y = y - 1
    | 'v' => y = y + 1
    end
    var pos = x.string() + "," + y.string()
    map = map.update(pos,1)

  be combine(p: Promise[Map[String,U32]]) =>
    p(map)

  
