local sysos = require "os"
local os = require "lualib/os"
local filepath = require "lualib/filepath"
local ioutil = require "lualib/ioutil"
local yaml = require "lualib/yaml"
local http = require "lualib/http"

if #arg ~= 1 then
    print("You must provide your `dotfiles.yml' as the only argument")
    sysos.exit(1)
end

local endpoint = "https://github.com/mdouchement/dotfiles/releases/download/binaries/"

local config_path = filepath.expand(arg[1])
local config = yaml.parse(ioutil.read_file(config_path))

for section, entries in pairs(config) do
    print("== "..section.." ==")

    for filename, entry in pairs(entries) do
        if entry["install"] then
            print("  Installing: "..filename)

            local src = filepath.join(section, filename)
            if section == "bin" then
                src = http.join(endpoint, filename.."-"..os.osname().."-amd64") -- limeline-darwin-amd64
            end

            local dst = filepath.expand(entry["path"]) -- Expand environment variables and custom characters

            os.mkdir_p(filepath.dirname(dst))            
            if section == "bin" then
                http.download(src, dst, true)
                os.chmod(dst, 0755)
            else
                os.cp_rf(src, dst)
            end
            
            print("  -> "..dst)
        end
    end

    print("\n")
end