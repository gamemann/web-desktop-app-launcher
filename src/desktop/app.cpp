#include <stdio.h>
#include <thread>
#include <chrono>

#include "include/cmdline.hpp"
#include "include/config.hpp"
#include "include/gui.hpp"

#define VERSION "1.0.0"

int main(int argc, char **argv) {
    int ret;

    // Parse command line.
    CmdLine cmd;
    
    if ((ret = ParseCmdLine(cmd, argc, argv)) != 0) {
        fprintf(stderr, "Error parsing command line. Return code => %d.\n", ret);

        return ret;
    }

    // Check for version.
    if (cmd.Version) {
        printf(VERSION);

        return 0;
    }

    // Check for help.
    if (cmd.Help) {
        PrintHelp();

        return 0;
    }

    // Parse config.
    Config cfg;
    
    if ((ret = ParseConfig(cfg, cmd.CfgPath)) != 0) {
        fprintf(stderr, "Error parsing config. Return code => %d.\n", ret);

        return ret;
    }

    // Check for list option.
    if (cmd.List) {
        ListConfig(cfg);

        return 0;
    }

    // Setup GUI application.
    SetupGui(cfg);

    while (true) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    return 0;
}