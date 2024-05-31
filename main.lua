local http_service = game:GetService('HttpService')

local console = {
	colors = {
		default = '[0m',
		black = '[30m',
		white = '[97m',
		
		red = '[31m',
		dark_green = '[32m',
		dark_yellow = '[33m',
		dark_blue = '[34m',
		dark_magenta = '[35m',
		dark_cyan = '[36m',
		dark_white = '[37m',
		bright_black = '[90m',
		bright_red = '[91m',
		bright_green = '[92m',
		bright_yellow = '[93m',
		bright_blue = '[94m',
		bright_magenta = '[95m',
		bright_cyan = '[96m'
	}
}

function console:initialize(host, port, internal)
	self.url = string.format('http://%s:%s', host, port)
	self.internal = internal
end

function console:fetch(data)
	if self.internal then
		local complete, result
		http_service:RequestInternal(data):Start(function(success, _result)
			result = success and _result
			complete = true
		end)
		repeat task.wait() until complete
		return result
	else
		return http_service:RequestAsync(data)
	end
end

function console:print(text)
	return self:fetch({
		Url = self.url .. '/print',
		Method = 'POST',
		Body = text
	}).StatusCode
end

function console:warn(text)
	return self:fetch({
		Url = self.url .. '/warn',
		Method = 'POST',
		Body = text
	}).StatusCode
end

function console:error(text)
	return self:fetch({
		Url = self.url .. '/error',
		Method = 'POST',
		Body = text
	}).StatusCode
end

function console:clear()
	return self:fetch({
		Url = self.url .. '/clear',
		Method = 'POST'
	}).StatusCode
end

function console:input(prompt)
	return self:fetch({
		Url = self.url .. '/input',
		Method = 'POST',
		Body = prompt
	}).Body
end

function console:title(title)
	return self:fetch({
		Url = self.url .. '/title',
		Method = 'POST',
		Body = title
	}).StatusCode
end

return console
