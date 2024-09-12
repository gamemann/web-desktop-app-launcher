#pragma once

#include <iostream>
#include <fstream>

#include <cerrno>
#include <cstring>

#include <nlohmann/json.hpp>

struct Config {
    int WindowWidth = 1920;
    int WindowHeight = 1080;
    bool FullScreen = true;
    std::string WebHost = "127.0.0.1";
    int WebPort = 2001;
} typedef Config;

static void ListConfig(Config& cfg) {
    printf("General\n");
    printf("\tWindow Width => %d.\n", cfg.WindowWidth);
    printf("\tWindow Height => %d.\n\n", cfg.WindowHeight);
    printf("Web Settings\n");
    printf("\tHost => %s.\n", cfg.WebHost.c_str());
    printf("\tPort => %d.\n", cfg.WebPort);
}

static int ParseConfig(Config& cfg, const std::string& path) {
    std::cout << "Parsing config file '" << path << "'..." << std::endl;
    
    std::ifstream f(path);

    if (!f.is_open()) {
        std::cerr << "Error opening config file :: " << std::strerror(errno) << std::endl;

        return 1;
    }

    // Parse JSON.
    nlohmann::json json_cfg;

    try {
        f >> json_cfg;
    } catch (nlohmann::json::parse_error& e) {
        std::cerr << "Error parsing JSON :: " << e.what() << std::endl;

        return 1;
    }

    // Assign values.
    try {
        cfg.WindowWidth = json_cfg.at("desktop").at("window_width").get<int>();
        cfg.WindowHeight = json_cfg.at("desktop").at("window_height").get<int>();
        cfg.FullScreen = json_cfg.at("desktop").at("full_screen").get<bool>();
        cfg.WebHost = json_cfg.at("web").at("host").get<std::string>();
        cfg.WebPort = json_cfg.at("web").at("port").get<int>();
    } catch (const nlohmann::json::exception& e) {
        std::cerr << "Error accessing and assigning JSON fields :: " << e.what() << std::endl;

        return 1;
    }

    return 0;
}