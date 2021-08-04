package ru.vladimirbalun.crashserpublisher.data.entity;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@EqualsAndHashCode
public class OSInfo {

    private String name = "";
    private String version = "";
    private String architecture = "";

}