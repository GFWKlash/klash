import QtQuick 2.15
import QtQuick.Controls 2.15
import Qt.labs.settings 1.0

import org.kde.kirigami 2.14 as Kirigami
import org.mauikit.controls 1.3 as Maui

Maui.ApplicationWindow
{
    id: root
    title: "Klash"

    headBar.visible: false

    property bool selectionMode: false
    property alias dialog :_dialogLoader.item

    Settings
    {
        id: viewerSettings
        property int viewType : Maui.AltBrowser.ViewType.Grid
    }

    Loader
    {
        id: _dialogLoader
    }

    StackView
    {
        id: _stackView
        anchors.fill: parent
    }
}
