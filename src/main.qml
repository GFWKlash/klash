import QtQuick 2.15
import QtQuick.Controls 2.15
import QtQuick.Layouts 1.3

import org.kde.kirigami 2.14 as Kirigami
import org.mauikit.controls 1.3 as Maui

import "views"

Maui.ApplicationWindow
{
    id: root
    title: "Klash"

    // Disable header bar
    headBar.visible: false

    sideBar: DesktopSidebar
    {
        id: desktopSidebar
    }

    StackView
    {
        id: _stackView
        anchors.fill: parent
        initialItem: MainView
        {
            id: mainView
            anchors.fill: parent
            showCSDControls: true

            headBar.leftContent: Maui.ToolButtonMenu
            {
                icon.name: "application-menu"

                MenuItem
                {
                    text: i18n("Home")
                    icon.name: "home"
                    onTriggered:
                    {
                        _stackViewTitle.text = i18n("Home")
                    }
                }

                MenuItem
                {
                    text: i18n("Data")
                    icon.name: "dashboard-show"
                    onTriggered:
                    {
                        _stackViewTitle.text = i18n("Data")
                    }
                }

                MenuItem
                {
                    text: i18n("Settings")
                    icon.name: "settings-configure"
                    onTriggered:
                    {
                        _stackViewTitle.text = i18n("Settings")
                    }
                }

                MenuItem
                {
                    text: i18n("About")
                    icon.name: "documentinfo"
                    onTriggered: root.about()
                }
            }

            headBar.rightContent: [
                // TODO: Run clash core button
                // TODO: Add profile button
                // TODO: Import profile button
            ]

            headBar.middleContent: RowLayout
            {
                spacing: 2
                Layout.fillWidth: true
                Layout.maximumWidth: 500
                Layout.minimumWidth: 0

                RowLayout
                {
                    Image {
                        sourceSize.height: Maui.Style.iconSizes.big
                        sourceSize.width: Maui.Style.iconSizes.big
                        source: "qrc:/assets/logo.png"
                    }

                    Text {
                        id: _stackViewTitle
                        text: i18n("Home")
                    }
                }
            }
        }
    }
}
