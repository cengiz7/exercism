defmodule Username do
  def sanitize(username) do
    Enum.map(username, fn c ->
      case c do
        ?ä -> 'ae'
        ?ö -> 'oe'
        ?ü -> 'ue'
        ?ß -> 'ss'
        ?_ -> '_'
        c when c >= ?a and c <= ?z -> c
        _ -> []
      end
    end)
    |> List.flatten()
  end
end
