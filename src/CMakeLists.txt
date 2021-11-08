# Copyright 2021 GFWK <gofuwako@gmail.com>
#
# SPDX-License-Identifier: GPL-3.0-or-later


set(klash_SRCS
    main.cpp
)

set(klash_ASSETS
    lib_assets.qrc
    )

if(QUICK_COMPILER)
    qtquick_compiler_add_resources(klash_QML_QRC qml.qrc)
else()
    qt5_add_resources(klash_QML_QRC qml.qrc)
endif()

if(ANDROID)
    add_library(${PROJECT_NAME} SHARED
        ${klash_SRCS}
        ${klash_ASSETS}
        ${klash_QML_QRC}
        )
else()
    add_executable(${PROJECT_NAME}
        ${klash_SRCS}
        ${klash_ASSETS}
        ${klash_QML_QRC}
        )
endif()

if (ANDROID)
     target_link_libraries(${PROJECT_NAME} Qt5::AndroidExtras Qt5::QuickControls2)
elseif(UNIX)
    target_link_libraries(${PROJECT_NAME} Qt5::Widgets)
endif()

target_compile_definitions(${PROJECT_NAME}
    PRIVATE $<$<OR:$<CONFIG:Debug>,$<CONFIG:RelWithDebInfo>>:QT_QML_DEBUG>)

target_link_libraries(${PROJECT_NAME} MauiKit MauiKit::FileBrowsing Qt5::Quick Qt5::Qml KF5::CoreAddons KF5::I18n)

if(UNIX AND NOT APPLE AND NOT ANDROID)
    install(TARGETS ${PROJECT_NAME} ${KDE_INSTALL_TARGETS_DEFAULT_ARGS})
    install(FILES assets/klash.svg DESTINATION ${KDE_INSTALL_ICONDIR}/hicolor/scalable/apps)
endif()
