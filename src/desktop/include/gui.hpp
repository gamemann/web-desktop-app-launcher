#pragma once

#include <iostream>

#include <gtk/gtk.h>
#include <webkit2/webkit2.h>

#include "config.hpp"

static void DestroyWindow(GtkWidget* widget, GtkWidget* window)
{
    gtk_main_quit();
}

static gboolean CloseWebView(WebKitWebView* webView, GtkWidget* window)
{
    gtk_widget_destroy(window);

    return true;
}

static int SetupGui(Config& cfg, int argc, char **argv) {
    // Format full web URL.
    char url[256];

    sprintf(url, "http://%s:%d", cfg.WebHost.c_str(), cfg.WebPort);

    std::cout << "Opening '" << url << "'!" << std::endl;

    // Init GTK.
    gtk_init(&argc, &argv);

    // Create main window with width/height from config.
    GtkWidget *main_window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
    gtk_window_set_default_size(GTK_WINDOW(main_window), cfg.WindowWidth, cfg.WindowHeight);

    // Check for full screen.
    if (cfg.FullScreen)
        gtk_window_fullscreen(GTK_WINDOW(main_window));

    // Create web view window.
    WebKitWebView *webView = WEBKIT_WEB_VIEW(webkit_web_view_new());

    gtk_container_add(GTK_CONTAINER(main_window), GTK_WIDGET(webView));

    // Create signals for destroying windows.
    g_signal_connect(main_window, "destroy", G_CALLBACK(DestroyWindow), NULL);
    g_signal_connect(webView, "close", G_CALLBACK(CloseWebView), main_window);

    // Load URL inside of web window.
    webkit_web_view_load_uri(webView, url);

    // Gain focus.
    gtk_widget_grab_focus(GTK_WIDGET(webView));

    // Show main window.
    gtk_widget_show_all(main_window);

    // Execute GTK.
    gtk_main();

    return 0;
}