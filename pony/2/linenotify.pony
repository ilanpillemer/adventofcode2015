use "buffered"

interface Processable 
  fun ref process(l: String)
  fun dispose()

class LineNotify is InputNotify
  let _rb: Reader
  let _p: Processable

  new create(p: Processable) =>
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
    try
      if _rb.size() > 0 then
        let rest: Array[U8] val = _rb.block(_rb.size())?
        _p.process(String.from_array(rest))
      end
    end
    _p.dispose()

