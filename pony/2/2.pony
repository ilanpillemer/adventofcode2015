use "regex"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(env, Process(env)) end, 512)

  class Process is Processable
    let  _env : Env
    new create(env: Env) =>
      _env = env
    
    fun process(l: String) =>
      _env.out.print(l)
      
   