--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: migrations_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.migrations_status AS ENUM (
    'update',
    'rollback'
);


ALTER TYPE public.migrations_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: adlinks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.adwords (
    id integer NOT NULL,
    result_id integer,
    type text NOT NULL,
    "position" text NOT NULL,
    link text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.adwords OWNER TO postgres;

--
-- Name: adwords_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.adwords_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.adwords_id_seq OWNER TO postgres;

--
-- Name: adwords_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.adwords_id_seq OWNED BY public.adwords.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.migrations (
    id_migration integer NOT NULL,
    name character varying(255) DEFAULT NULL::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    statements text,
    rollback_statements text,
    status public.migrations_status
);


ALTER TABLE public.migrations OWNER TO postgres;

--
-- Name: migrations_id_migration_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.migrations_id_migration_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.migrations_id_migration_seq OWNER TO postgres;

--
-- Name: migrations_id_migration_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.migrations_id_migration_seq OWNED BY public.migrations.id_migration;


--
-- Name: oauth2_clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oauth2_clients (
    id text NOT NULL,
    secret text NOT NULL,
    domain text NOT NULL,
    data jsonb NOT NULL
);


ALTER TABLE public.oauth2_clients OWNER TO postgres;

--
-- Name: oauth2_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oauth2_tokens (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    code text NOT NULL,
    access text NOT NULL,
    refresh text NOT NULL,
    data jsonb NOT NULL
);


ALTER TABLE public.oauth2_tokens OWNER TO postgres;

--
-- Name: oauth2_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.oauth2_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.oauth2_tokens_id_seq OWNER TO postgres;

--
-- Name: oauth2_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.oauth2_tokens_id_seq OWNED BY public.oauth2_tokens.id;


--
-- Name: results; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.results (
    id integer NOT NULL,
    user_id integer,
    keyword text NOT NULL,
    status text NOT NULL,
    non_ad_links json,
    page_cache text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.results OWNER TO postgres;

--
-- Name: results_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.results_id_seq OWNER TO postgres;

--
-- Name: results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.results_id_seq OWNED BY public.results.id;


--
-- Name: session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.session (
    session_key character(64) NOT NULL,
    session_data bytea,
    session_expiry timestamp without time zone NOT NULL
);


ALTER TABLE public.session OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email public.citext,
    hashed_password text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.users.id;


--
-- Name: adlinks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adwords ALTER COLUMN id SET DEFAULT nextval('public.adwords_id_seq'::regclass);


--
-- Name: migrations id_migration; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id_migration SET DEFAULT nextval('public.migrations_id_migration_seq'::regclass);


--
-- Name: oauth2_tokens id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth2_tokens ALTER COLUMN id SET DEFAULT nextval('public.oauth2_tokens_id_seq'::regclass);


--
-- Name: results id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.results ALTER COLUMN id SET DEFAULT nextval('public.results_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Name: adlinks adword_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adwords
    ADD CONSTRAINT adword_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id_migration);


--
-- Name: oauth2_clients oauth2_clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth2_clients
    ADD CONSTRAINT oauth2_clients_pkey PRIMARY KEY (id);


--
-- Name: oauth2_tokens oauth2_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth2_tokens
    ADD CONSTRAINT oauth2_tokens_pkey PRIMARY KEY (id);


--
-- Name: results result_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.results
    ADD CONSTRAINT result_pkey PRIMARY KEY (id);


--
-- Name: session session_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.session
    ADD CONSTRAINT session_key PRIMARY KEY (session_key);


--
-- Name: users user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: idx_oauth2_tokens_access; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oauth2_tokens_access ON public.oauth2_tokens USING btree (access);


--
-- Name: idx_oauth2_tokens_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oauth2_tokens_code ON public.oauth2_tokens USING btree (code);


--
-- Name: idx_oauth2_tokens_expires_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oauth2_tokens_expires_at ON public.oauth2_tokens USING btree (expires_at);


--
-- Name: idx_oauth2_tokens_refresh; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oauth2_tokens_refresh ON public.oauth2_tokens USING btree (refresh);


--
-- Name: adlinks adwords_result_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adwords
    ADD CONSTRAINT adwords_result_id_fkey FOREIGN KEY (result_id) REFERENCES public.results(id) ON DELETE CASCADE;


--
-- Name: results results_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.results
    ADD CONSTRAINT results_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

