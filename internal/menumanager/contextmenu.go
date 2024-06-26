package menumanager

import (
	"encoding/json"
	"fmt"

	"github.com/888go/wails/pkg/menu"
)

type ContextMenu struct {
	ID            string
	ProcessedMenu *WailsMenu
	menuItemMap   *MenuItemMap
	menu          *menu.Menu
}


// ff:
func (t *ContextMenu) AsJSON() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


// ff:
// contextMenu:
func NewContextMenu(contextMenu *menu.ContextMenu) *ContextMenu {
	result := &ContextMenu{
		ID:          contextMenu.ID,
		menu:        contextMenu.X菜单,
		menuItemMap: NewMenuItemMap(),
	}

	result.menuItemMap.AddMenu(contextMenu.X菜单)
	result.ProcessedMenu = NewWailsMenu(result.menuItemMap, result.menu)

	return result
}


// ff:
// contextMenu:
func (m *Manager) AddContextMenu(contextMenu *menu.ContextMenu) {
	newContextMenu := NewContextMenu(contextMenu)

	// Save the references
	m.contextMenus[contextMenu.ID] = newContextMenu
	m.contextMenuPointers[contextMenu] = contextMenu.ID
}


// ff:
// contextMenu:
func (m *Manager) UpdateContextMenu(contextMenu *menu.ContextMenu) (string, error) {
	contextMenuID, contextMenuKnown := m.contextMenuPointers[contextMenu]
	if !contextMenuKnown {
		return "", fmt.Errorf("unknown Context Menu '%s'. Please add the context menu using AddContextMenu()", contextMenu.ID)
	}

	// 创建更新后的上下文菜单
	updatedContextMenu := NewContextMenu(contextMenu)

	// Save the reference
	m.contextMenus[contextMenuID] = updatedContextMenu

	return updatedContextMenu.AsJSON()
}
