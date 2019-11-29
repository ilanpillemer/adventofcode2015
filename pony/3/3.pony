use "debug"
use "collections"

actor Main
  new create(env: Env) =>
    env.out.print("day3")
    let collector = Collector
    env.input(Director(collector),512)
    

class iso Director is InputNotify
 let santa:Santa
 let robo:Santa
 var turn:Bool = false

 new iso create(collector: Collector) =>
  santa = Santa(collector)
  collector.register()
  robo = Santa(collector)
  collector.register()

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
    santa.dispose()
    robo.dispose()

actor Santa
  var x: I32 = 0
  var y: I32 = 0
  let _collector:Collector tag
  
  var map: Map[String,U32] iso = recover Map[String,U32] end

  new create(collector:Collector) =>
    _collector = collector
    map.insert("0,0",1)

  be move(d: U8) =>
    match d
    | '<' => x = x - 1
    | '>' => x = x + 1
    | '^' => y = y - 1
    | 'v' => y = y + 1
    end
    var pos = x.string() + "," + y.string()
    map.insert(pos,1)

  be dispose() =>
    var temp: Map[String,U32] iso = map = recover Map[String,U32] end
    _collector.update_map(consume temp)
    _collector.dispose()    

  actor Collector
    var count:U8 = 0
    var map:Map[String,I32] = Map[String,I32]

  be register() => count = count + 1

  be update_map(map': Map[String,U32] iso) =>
    var m: Map[String,U32] ref = consume map'
    for k in m.keys() do
     map.insert(k,0)
    end

  be dispose() =>
   count = count - 1
   if count == 0 then Debug.out("collected results: " + map.size().string()) end
