package ru.vladimirbalun.crashserpublisher.data.entity;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.util.UUID;

@Getter
@Setter
@Entity
@EqualsAndHashCode
@Table(name = "app_info")
public class AppInfo {

    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    @Column(name = "id_app_info", columnDefinition = "BINARY(16)")
    private UUID id_app_info;

    @Column(name="name", nullable = true, length = 250)
    private String name = "";

    @Column(name="username", nullable = true, length = 50)
    private String version = "";

    @ManyToOne
    @JoinColumn(name="programming_language", nullable=false)
    private ProgrammingLanguage programmingLanguage;

}
