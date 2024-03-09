--
-- PostgreSQL database dump
--

-- Dumped from database version 11.22
-- Dumped by pg_dump version 11.22

-- Started on 2024-03-09 19:49:58

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

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 32781)
-- Name: orders; Type: TABLE; Schema: public; Owner: l0user
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    uid character varying(50) NOT NULL,
    orderjson jsonb NOT NULL
);


ALTER TABLE public.orders OWNER TO l0user;

--
-- TOC entry 196 (class 1259 OID 32779)
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: l0user
--

CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO l0user;

--
-- TOC entry 2817 (class 0 OID 0)
-- Dependencies: 196
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: l0user
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- TOC entry 2686 (class 2604 OID 32784)
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: l0user
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- TOC entry 2811 (class 0 OID 32781)
-- Dependencies: 197
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: l0user
--

COPY public.orders (id, uid, orderjson) FROM stdin;
\.


--
-- TOC entry 2818 (class 0 OID 0)
-- Dependencies: 196
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: l0user
--

SELECT pg_catalog.setval('public.orders_id_seq', 5, true);


--
-- TOC entry 2688 (class 2606 OID 32789)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: l0user
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


-- Completed on 2024-03-09 19:49:58

--
-- PostgreSQL database dump complete
--

