// Copyright 2018-2020 Camilo Higuita <milo.h@aol.com>
// Copyright 2018-2020 Nitrux Latinoamericana S.C.
//
// SPDX-License-Identifier: GPL-3.0-or-later


import QtQuick 2.14
import QtQml 2.14

import QtQuick.Controls 2.14

import org.mauikit.controls 1.3 as Maui

import org.kde.kirigami 2.6 as Kirigami

Maui.AbstractSideBar
{
    id: control

    collapsible: true
    collapsed : !root.isWide
    preferredWidth: Kirigami.Units.gridUnit
        * (Maui.Handy.isWindows ?  15 : 13)

    ListModel {
        id: desktopSidebarModel

        ListElement {
            label: "Home"
            icon: "home"
        }
        ListElement {
            label: "Data"
            icon: "dashboard-show"
        }
        ListElement {
            label: "Settings"
            icon: "settings-configure"
        }
    }

    Maui.ListBrowser
    {
        anchors.fill: parent
        id: _listBrowser
        topPadding: 0
        bottomPadding: 0

        flickable.topMargin: Maui.Style.space.medium
        flickable.bottomMargin: Maui.Style.space.medium
        flickable.header: Column
        {
            width: parent.width
            spacing: Maui.Style.space.medium

            GridView
            {
                id: _quickSection
                implicitHeight: contentHeight + Maui.Style.space.medium * 1.5
                currentIndex : 0
                width: parent.width
                cellWidth: Math.floor(parent.width)
                cellHeight: 80
                interactive: false

                model: desktopSidebarModel

                delegate: Item
                {
                    height: GridView.view.cellHeight
                    width: GridView.view.cellWidth

                    Maui.GridBrowserDelegate
                    {
                        radius: 15
                        isCurrentItem: parent.GridView.isCurrentItem
                        anchors.fill: parent
                        anchors.margins: Maui.Style.space.tiny
                        iconSource: model.icon +  (Qt.platform.os == "android" || Qt.platform.os == "osx" ? ("-sidebar") : "")
                        iconSizeHint: Maui.Style.iconSizes.huge
                        // template.isMask: true
                        label1.text: model.label
                        labelsVisible: true
                        tooltipText: model.label
                        onClicked:
                        {
                            console.log("Item clicked")
                        }
                    }
                }
            }
        }
    }
}
