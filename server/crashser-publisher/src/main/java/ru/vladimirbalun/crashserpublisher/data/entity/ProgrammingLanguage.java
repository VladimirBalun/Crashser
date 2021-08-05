package ru.vladimirbalun.crashserpublisher.data.entity;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Getter
@Setter
@Entity
@EqualsAndHashCode
@Table(name = "programming_language")
public class ProgrammingLanguage {

    @Id
    @Column(name="name", unique = true, nullable = false, length = 50)
    private String name;

}
