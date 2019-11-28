use "buffered"

interface Processable
  fun process(l: String)

class LineNotify
  let _env: Env
  var _rb: Reader
  var _p: Processable

  new create(env: Env, p: Processable ) =>
    _env = env
    _rb = Reader
    _p = p

  fun ref apply(data: Array[U8] iso) =>
    _rb.append(consume data)
    while true do
      try
        let l = _rb.line()?
	_p.process(consume l)
      else
        break
      end
    end

  fun ref dispose() =>
    None


