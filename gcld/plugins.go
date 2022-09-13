package gcld

import (
	"github.com/robertkrimen/otto"
	"os"
	"path/filepath"
	"sockets-proxy/util/log"
)

type PluginManager struct {
	pluginFolder string
	plugins      map[string]*Plugin
}
type Plugin struct {
	Name    string
	Command string

	Vm *otto.Otto
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func NewPluginManager(pluginFolder string) *PluginManager {
	pluginManager := &PluginManager{pluginFolder: pluginFolder}
	pluginManager.plugins = make(map[string]*Plugin)
	pluginManager.Load()
	return pluginManager
}
func (p *PluginManager) Load() {

	err := filepath.Walk(p.pluginFolder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		plugin := NewPlugin()
		data, _ := os.ReadFile(path)
		vm := otto.New()
		vm.Run(data)
		name, err := vm.Call("Plugin.Name", nil)
		if err != nil {
			log.Log.Errorf("load plugin %s error:%s", path, err.Error())
			return nil
		}
		plugin.Name = name.String()

		command, err := vm.Call("Plugin.Command", nil)
		if err != nil {
			log.Log.Errorf("load plugin %s error:%s", path, err.Error())
			return nil
		}
		plugin.Command = command.String()
		log.Log.Infof("plugin %s->%s", plugin.Command, plugin.Name)
		plugin.Vm = vm
		p.plugins[plugin.Command] = plugin
		return nil
	})
	if err != nil {
		panic(err)
	}

}
