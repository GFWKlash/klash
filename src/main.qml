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

    property bool selectionMode: false

    MainView
    {
        id: mainView
        anchors.fill: parent
        showCSDControls: true

        headBar.forceCenterMiddleContent: false

        headBar.leftContent: Maui.ToolButtonMenu
        {
            icon.name: "application-menu"

            MenuItem
            {
                text: i18n("Settings")
                icon.name: "settings-configure"
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
            ToolButton
            {
                icon.name: "list-add"
                onClicked:
                {
                }
            },
            // TODO: Add profile button
            ToolButton
            {
                icon.name: "list-add"
                onClicked:
                {
                }
            },
            // TODO: Import profile button
            ToolButton
            {
                icon.name: "list-add"
                onClicked:
                {
                }
            }
        ]

        headBar.middleContent: RowLayout
        {
            anchors.fill: parent
            spacing: 2
            Layout.fillWidth: true
            Layout.maximumWidth: 500
            Layout.minimumWidth: 0

            Text {
                text: i18n("Profiles")
            }
        }
    }
}
