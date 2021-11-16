import QtQuick 2.14
import QtQuick.Controls 2.13

import org.kde.kirigami 2.14 as Kirigami

import org.mauikit.controls 1.3 as Maui

import QtQuick.Window 2.0

Maui.Page
{
    id: control

    // Show headbar when the sidebar is collapsed
    headBar.visible: !root.isWide

    Text {
        text: "Hello Klash"
    }

}
