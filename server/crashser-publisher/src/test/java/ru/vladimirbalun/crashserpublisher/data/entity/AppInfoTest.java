package ru.vladimirbalun.crashserpublisher.data.entity;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AppInfoTest {

    @Test
    void whenAppInfoIsEmptyThenReturnEmptyFields() {
        final AppInfo appInfo = new AppInfo();
        assertThat("").isEqualTo(appInfo.getName());
        assertThat("").isEqualTo(appInfo.getVersion());
        assertThat("").isEqualTo(appInfo.getProgrammingLanguage());
    }

    @Test
    void whenSetValidNameThenReturnValidName() {
        final String name = "test_name";
        AppInfo appInfo = new AppInfo();
        appInfo.setName(name);
        assertThat(name).isEqualTo(appInfo.getName());
    }

    @Test
    void whenSetValidVersionThenReturnValidVersion() {
        final String version = "1.0.0";
        AppInfo appInfo = new AppInfo();
        appInfo.setVersion(version);
        assertThat(version).isEqualTo(appInfo.getVersion());
    }

    @Test
    void whenSetValidProgrammingLanguageThenReturnValidProgrammingLanguage() {
        final String programmingLanguage = "C++";
        AppInfo appInfo = new AppInfo();
        appInfo.setProgrammingLanguage(programmingLanguage);
        assertThat(programmingLanguage).isEqualTo(appInfo.getProgrammingLanguage());
    }

}
