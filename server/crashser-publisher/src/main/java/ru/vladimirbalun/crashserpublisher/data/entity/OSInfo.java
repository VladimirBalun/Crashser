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
@Table(name = "os_info")
public class OSInfo {

    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    @Column(name = "id_os_info", columnDefinition = "BINARY(16)")
    private UUID id_os_info;

    @Column(name="name", nullable = true, length = 250)
    private String name;

    @Column(name="version", nullable = true, length = 250)
    private String version;

}