#pragma once

#include <iostream>
#include <getopt.h>

struct cmdline {
    std::string CfgPath = "/etc/wdal/conf.json";
    bool List = false;
    bool Version = false;
    bool Help = false;
} typedef CmdLine;

static struct option longOpts[] = {
    { "cfgpath", required_argument, nullptr, 'c' },
    { "list", no_argument, nullptr, 'l' },
    { "version", no_argument, nullptr, 'v' },
    { "help", no_argument, nullptr, 'h' },
    { nullptr, 0, nullptr, 0 }
};

static void PrintHelp() {
    printf("Usage: wdal-app -c <path> -l -v -h\n");
    printf("-c --cfgpath => The path to the config file (default /etc/wdal/conf.json).\n");
    printf("-l --list => Lists config options.\n");
    printf("-v --version => Prints the current version.\n");
    printf("-h --help => Prints the help menu.\n");
}

static int ParseCmdLine(CmdLine& cmd, int argc, char** argv) {
    int opt;

    while ((opt = getopt_long(argc, argv, "c:lvh", longOpts, nullptr)) != -1) {
        switch (opt) {
            case 'c':
                cmd.CfgPath = optarg;

                break; 
            
            case 'l':
                cmd.List = true;

                break;

            case 'v':
                cmd.Version = true;

                break;

            case 'h':
                cmd.Help = true;

                break;

            case '?':
                cmd.Help = true;

                break;

            default:
                std::cerr << "Unknown option: " << opt << std::endl;

                return 1;
        }
    }

    return 0;
}