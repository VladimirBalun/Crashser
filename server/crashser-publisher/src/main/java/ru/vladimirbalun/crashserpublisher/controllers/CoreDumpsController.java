package ru.vladimirbalun.crashserpublisher.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CoreDumpsController {

    @GetMapping(value = "/core_dump_list/{page_number}", produces = "application/json; charset=utf-8")
    public String getCoreDumpList(@PathVariable("page_number") long pageNumber) {
        return "Page";
    }

}
