CREATE DATABASE flights
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.locations

-- DROP TABLE public.locations;

CREATE TABLE public.locations
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name character varying COLLATE pg_catalog."default",
    CONSTRAINT location_pkey PRIMARY KEY (id),
    CONSTRAINT name_uq UNIQUE (name)

)

TABLESPACE pg_default;

ALTER TABLE public.locations
    OWNER to postgres;

-- Index: location_idx

-- DROP INDEX public.location_idx;

CREATE INDEX location_idx
    ON public.locations USING btree
    (id)
    TABLESPACE pg_default;


    -- Table: public.reservations

-- DROP TABLE public.reservations;

CREATE TABLE public.reservations
(
    date timestamp without time zone NOT NULL,
    reservation character varying COLLATE pg_catalog."default" NOT NULL,
    location_id bigint,
    CONSTRAINT location_id_fk FOREIGN KEY (location_id)
        REFERENCES public.locations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.reservations
    OWNER to postgres;

-- Index: fki_location_id_fk

-- DROP INDEX public.fki_location_id_fk;

CREATE INDEX fki_location_id_fk
    ON public.reservations USING btree
    (location_id)
    TABLESPACE pg_default;
