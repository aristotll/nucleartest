local i = 10
local seq = {}

function random()
	while (i > 0) do 
		seq[i] = math.random(i)
		i = i - 1
	end

	return seq
end

function printTable(table)
	for k, v in ipairs(table) do
		print("[" .. k .. "]" .. "=>" .. v)
	end
end

print(printTable(random()))
