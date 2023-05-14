--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

-- Started on 2023-05-10 14:28:54

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
-- TOC entry 247 (class 1255 OID 50815)
-- Name: delete_emplpost(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_emplpost(p_employee_id integer) RETURNS character
    LANGUAGE plpgsql
    AS $$
begin
	delete from Employee_post where Employee_ID=p_Employee_ID;
	return 'deleted';
end;
$$;


ALTER FUNCTION public.delete_emplpost(p_employee_id integer) OWNER TO postgres;

--
-- TOC entry 246 (class 1255 OID 50814)
-- Name: delete_emplpost(integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_emplpost(p_post_id integer, p_employee_id integer) RETURNS character
    LANGUAGE plpgsql
    AS $$
begin
	delete from Employee_post where Post_ID = p_Post_ID AND Employee_ID=p_Employee_ID;
	return 'deleted';
end;
$$;


ALTER FUNCTION public.delete_emplpost(p_post_id integer, p_employee_id integer) OWNER TO postgres;

--
-- TOC entry 251 (class 1255 OID 50881)
-- Name: delete_logrow(character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_logrow(tablename character varying, primary_colname character varying, col_value integer) RETURNS character
    LANGUAGE plpgsql
    AS $_$
begin
	execute format('UPDATE %s SET log_del=true where %s = $1', tablename, primary_colname)
	using col_value;
	return 'deleted';
end;
$_$;


ALTER FUNCTION public.delete_logrow(tablename character varying, primary_colname character varying, col_value integer) OWNER TO postgres;

--
-- TOC entry 232 (class 1255 OID 50733)
-- Name: delete_row(character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_row(tablename character varying, table_colname character varying, col_value integer) RETURNS character
    LANGUAGE plpgsql
    AS $_$
begin
	execute format('delete from %s where %s = $1', tablename, table_colname)
	using col_value;
	return 'deleted';
end;
$_$;


ALTER FUNCTION public.delete_row(tablename character varying, table_colname character varying, col_value integer) OWNER TO postgres;

--
-- TOC entry 223 (class 1255 OID 50719)
-- Name: generate_key(integer, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.generate_key(p_id integer, p_auth_key character varying) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update organization
	set
	auth_key = $2
	where id_organization = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.generate_key(p_id integer, p_auth_key character varying) OWNER TO postgres;

--
-- TOC entry 231 (class 1255 OID 50732)
-- Name: insert_department(character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_department(p_name character varying, p_description character varying, p_organization_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into Department(Name,Description,Organization_ID) 
	values (p_Name,p_Description,p_Organization_ID) returning id_department into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_department(p_name character varying, p_description character varying, p_organization_id integer) OWNER TO postgres;

--
-- TOC entry 244 (class 1255 OID 50756)
-- Name: insert_employee(character varying, character varying, character varying, date, character varying, character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_employee(p_surname character varying, p_name character varying, p_secondname character varying, p_date_birth date, p_seriapasp character varying, p_numberpasp character varying, p_email character varying, p_department_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into Employee(Surname,Name,SecondName,Date_Birth,SeriaPasp, NumberPasp, Email, Department_ID) 
	values (p_Surname,p_Name,p_SecondName,p_Date_Birth,p_SeriaPasp, p_NumberPasp, p_Email, p_Department_ID) returning id_employee into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_employee(p_surname character varying, p_name character varying, p_secondname character varying, p_date_birth date, p_seriapasp character varying, p_numberpasp character varying, p_email character varying, p_department_id integer) OWNER TO postgres;

--
-- TOC entry 227 (class 1255 OID 50728)
-- Name: insert_employee_post(integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_employee_post(p_post_id integer, p_employee_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into Employee_Post(Post_ID,Employee_ID) 
	values (p_Post_ID,p_Employee_ID) returning id_employee_post into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_employee_post(p_post_id integer, p_employee_id integer) OWNER TO postgres;

--
-- TOC entry 226 (class 1255 OID 50722)
-- Name: insert_finances_operations(numeric, date, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_finances_operations(p_summ numeric, p_date_operation date, p_description character varying, p_organization_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into Finances_Operations(Summ,Date_Operation,Description,Organization_ID) 
	values (p_Summ,p_Date_Operation,p_Description,p_Organization_ID) returning id_operations into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_finances_operations(p_summ numeric, p_date_operation date, p_description character varying, p_organization_id integer) OWNER TO postgres;

--
-- TOC entry 224 (class 1255 OID 50720)
-- Name: insert_organization(character varying, character varying, character varying, numeric, date, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_organization(p_name character varying, p_addres character varying, p_inn character varying, p_budget numeric, p_date_foundation date, p_auth_key character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into organization(name,addres,inn,budget,date_foundation, auth_key) 
	values (p_name,p_addres,p_inn,p_budget, p_date_foundation, p_Auth_Key) returning id_organization into tempid;
	RETURN tempid;
end;
$$;


ALTER FUNCTION public.insert_organization(p_name character varying, p_addres character varying, p_inn character varying, p_budget numeric, p_date_foundation date, p_auth_key character varying) OWNER TO postgres;

--
-- TOC entry 229 (class 1255 OID 50730)
-- Name: insert_post(character varying, numeric, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_post(p_name character varying, p_salary numeric, p_department_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into Post(Name,Salary,Department_ID) 
	values (p_Name,p_Salary,p_Department_ID) returning id_post into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_post(p_name character varying, p_salary numeric, p_department_id integer) OWNER TO postgres;

--
-- TOC entry 253 (class 1255 OID 50875)
-- Name: insert_sgt(character varying, character varying, date, date, boolean, timestamp without time zone, integer, character varying, character varying, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_sgt(p_name character varying, p_descr character varying, p_datetstart date, p_dateend date, p_done boolean, p_date_done timestamp without time zone, p_foreignkey integer, foreignkeyname character varying, tablename character varying, prkeyname character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $_$
DECLARE tempid int;
begin
	execute format('insert into %s(name,Description,Date_Start,Date_End, done, date_done, %s) 
	values ($1,$2,$3,$4, $5, $6, $7) 
				   returning %s', tablename, foreignkeyname, prkeyname) into tempid
	using $1,$2,$3,$4, $5, $6, $7;
	return tempid;
end;
$_$;


ALTER FUNCTION public.insert_sgt(p_name character varying, p_descr character varying, p_datetstart date, p_dateend date, p_done boolean, p_date_done timestamp without time zone, p_foreignkey integer, foreignkeyname character varying, tablename character varying, prkeyname character varying) OWNER TO postgres;

--
-- TOC entry 250 (class 1255 OID 50754)
-- Name: insert_user(character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_user(p_login character varying, p_password character varying, p_employee_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE tempid int;
begin
	insert into "User"(Login,Password, Employee_ID) 
	values (p_Login,p_Password,p_Employee_ID) returning id_user into tempid;
	return tempid;
end;
$$;


ALTER FUNCTION public.insert_user(p_login character varying, p_password character varying, p_employee_id integer) OWNER TO postgres;

--
-- TOC entry 252 (class 1255 OID 50885)
-- Name: swap_roles(integer, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.swap_roles(p_olduser integer, p_newuser integer, p_oldrole integer, p_newrole integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update "User"
	set
	roles_id=$4
	where employee_id = $2;
	update "User"
	set
	roles_id=$3
	where employee_id = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.swap_roles(p_olduser integer, p_newuser integer, p_oldrole integer, p_newrole integer) OWNER TO postgres;

--
-- TOC entry 230 (class 1255 OID 50731)
-- Name: update_department(integer, character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_department(p_id integer, p_name character varying, p_description character varying, p_organization_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update Department
	set
	Name=$2,
	Description = $3,
	Organization_ID=$4
	where ID_Department = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_department(p_id integer, p_name character varying, p_description character varying, p_organization_id integer) OWNER TO postgres;

--
-- TOC entry 249 (class 1255 OID 50856)
-- Name: update_emailverification(character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_emailverification(p_email character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $_$
DECLARE tempid int;
begin
execute format('
	update Employee
	set
	EmailVerified=true
	where Email = $1
	returning ID_Employee') into tempid
	using $1;
	return tempid;
end;
$_$;


ALTER FUNCTION public.update_emailverification(p_email character varying) OWNER TO postgres;

--
-- TOC entry 245 (class 1255 OID 50757)
-- Name: update_employee(integer, character varying, character varying, character varying, date, character varying, character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_employee(p_id integer, p_surname character varying, p_name character varying, p_secondname character varying, p_date_birth date, p_seriapasp character varying, p_numberpasp character varying, p_email character varying, p_department_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update Employee
	set
	Surname = $2,
	Name = $3,
	SecondName = $4,
	Date_Birth=$5,
	SeriaPasp=$6,
	NumberPasp=$7,
	Email=$8,
	Department_ID=$9
	where ID_Employee = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_employee(p_id integer, p_surname character varying, p_name character varying, p_secondname character varying, p_date_birth date, p_seriapasp character varying, p_numberpasp character varying, p_email character varying, p_department_id integer) OWNER TO postgres;

--
-- TOC entry 225 (class 1255 OID 50721)
-- Name: update_finances_operations(integer, numeric, date, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_finances_operations(p_id integer, p_summ numeric, p_date_operation date, p_description character varying, p_organization_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update Finances_Operations
	set
	Summ = $2,
	Date_Operation = $3,
	Description = $4,
	Organization_ID=$5
	where ID_Operations = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_finances_operations(p_id integer, p_summ numeric, p_date_operation date, p_description character varying, p_organization_id integer) OWNER TO postgres;

--
-- TOC entry 222 (class 1255 OID 50718)
-- Name: update_organization(integer, character varying, character varying, character varying, numeric, date); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_organization(p_id integer, p_name character varying, p_addres character varying, p_inn character varying, p_budget numeric, p_date_foundation date) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update organization
	set
	name = $2,
	addres = $3,
	inn = $4,
	budget = $5,
	date_foundation = $6
	where id_organization = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_organization(p_id integer, p_name character varying, p_addres character varying, p_inn character varying, p_budget numeric, p_date_foundation date) OWNER TO postgres;

--
-- TOC entry 228 (class 1255 OID 50729)
-- Name: update_post(integer, character varying, numeric, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_post(p_id integer, p_name character varying, p_salary numeric, p_department_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update Post
	set
	Name = $2,
	Salary = $3,
	Department_ID=$4
	where ID_Post = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_post(p_id integer, p_name character varying, p_salary numeric, p_department_id integer) OWNER TO postgres;

--
-- TOC entry 254 (class 1255 OID 50876)
-- Name: update_sgt(integer, character varying, character varying, date, date, boolean, timestamp without time zone, integer, character varying, character varying, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_sgt(p_id integer, p_name character varying, p_descr character varying, p_datetstart date, p_dateend date, p_done boolean, p_date_done timestamp without time zone, p_foreignkey integer, foreignkeyname character varying, tablename character varying, prkeyname character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $_$
DECLARE tempid int;
begin
execute format('
	update %s
	set
	name = $2,
	Description = $3,
	Date_start = $4,
	Date_end = $5,
	Done = $6,
	date_done = $7,
	%s=$8
	where %s = $1;', tablename, foreignkeyname, prkeyname)
	using $1, $2, $3, $4, $5, $6, $7, $8;
	return tempid;
end;
$_$;


ALTER FUNCTION public.update_sgt(p_id integer, p_name character varying, p_descr character varying, p_datetstart date, p_dateend date, p_done boolean, p_date_done timestamp without time zone, p_foreignkey integer, foreignkeyname character varying, tablename character varying, prkeyname character varying) OWNER TO postgres;

--
-- TOC entry 248 (class 1255 OID 50755)
-- Name: update_user(integer, character varying, character varying, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_user(p_id integer, p_login character varying, p_password character varying, p_employee_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
begin
	update "User"
	set
	Login = $2,
	Password = $3,
	Employee_ID=$4
	where ID_User = $1;
	return true;
end;
$_$;


ALTER FUNCTION public.update_user(p_id integer, p_login character varying, p_password character varying, p_employee_id integer) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 205 (class 1259 OID 50521)
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    id_user integer NOT NULL,
    login character varying(50) NOT NULL,
    password character varying(250) NOT NULL,
    employee_id integer,
    roles_id integer DEFAULT 1
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 50519)
-- Name: User_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."User_id_user_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."User_id_user_seq" OWNER TO postgres;

--
-- TOC entry 3168 (class 0 OID 0)
-- Dependencies: 204
-- Name: User_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."User_id_user_seq" OWNED BY public."User".id_user;


--
-- TOC entry 203 (class 1259 OID 50507)
-- Name: department; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.department (
    id_department integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(250) NOT NULL,
    organization_id integer NOT NULL
);


ALTER TABLE public.department OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 50505)
-- Name: department_id_department_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.department_id_department_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.department_id_department_seq OWNER TO postgres;

--
-- TOC entry 3169 (class 0 OID 0)
-- Dependencies: 202
-- Name: department_id_department_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.department_id_department_seq OWNED BY public.department.id_department;


--
-- TOC entry 215 (class 1259 OID 50665)
-- Name: employee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employee (
    id_employee integer NOT NULL,
    surname character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    secondname character varying(50) DEFAULT ' '::character varying,
    date_birth date NOT NULL,
    seriapasp character varying(4) NOT NULL,
    numberpasp character varying(6) NOT NULL,
    email character varying(50) NOT NULL,
    department_id integer,
    emailverified boolean DEFAULT false
);


ALTER TABLE public.employee OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 50663)
-- Name: employee_id_employee_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employee_id_employee_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employee_id_employee_seq OWNER TO postgres;

--
-- TOC entry 3170 (class 0 OID 0)
-- Dependencies: 214
-- Name: employee_id_employee_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employee_id_employee_seq OWNED BY public.employee.id_employee;


--
-- TOC entry 217 (class 1259 OID 50685)
-- Name: employee_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employee_post (
    id_employee_post integer NOT NULL,
    post_id integer NOT NULL,
    employee_id integer NOT NULL
);


ALTER TABLE public.employee_post OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 50683)
-- Name: employee_post_id_employee_post_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employee_post_id_employee_post_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employee_post_id_employee_post_seq OWNER TO postgres;

--
-- TOC entry 3171 (class 0 OID 0)
-- Dependencies: 216
-- Name: employee_post_id_employee_post_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employee_post_id_employee_post_seq OWNED BY public.employee_post.id_employee_post;


--
-- TOC entry 207 (class 1259 OID 50607)
-- Name: finances_operations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.finances_operations (
    id_operations integer NOT NULL,
    summ numeric(36,2) NOT NULL,
    date_operation date NOT NULL,
    description character varying(250) NOT NULL,
    organization_id integer NOT NULL
);


ALTER TABLE public.finances_operations OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 50605)
-- Name: finances_operations_id_operations_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.finances_operations_id_operations_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.finances_operations_id_operations_seq OWNER TO postgres;

--
-- TOC entry 3172 (class 0 OID 0)
-- Dependencies: 206
-- Name: finances_operations_id_operations_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.finances_operations_id_operations_seq OWNED BY public.finances_operations.id_operations;


--
-- TOC entry 209 (class 1259 OID 50621)
-- Name: goal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goal (
    id_goal integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(250) NOT NULL,
    date_start date NOT NULL,
    date_end date NOT NULL,
    done boolean DEFAULT false,
    department_id integer NOT NULL,
    date_done timestamp without time zone,
    log_del boolean DEFAULT false
);


ALTER TABLE public.goal OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 50619)
-- Name: goal_id_goal_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goal_id_goal_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goal_id_goal_seq OWNER TO postgres;

--
-- TOC entry 3173 (class 0 OID 0)
-- Dependencies: 208
-- Name: goal_id_goal_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.goal_id_goal_seq OWNED BY public.goal.id_goal;


--
-- TOC entry 201 (class 1259 OID 50495)
-- Name: organization; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organization (
    id_organization integer NOT NULL,
    name character varying(250) NOT NULL,
    addres character varying(250) NOT NULL,
    inn character varying(13) NOT NULL,
    budget numeric(36,2) NOT NULL,
    auth_key character varying(255) DEFAULT 'NoKey'::character varying,
    date_foundation date NOT NULL
);


ALTER TABLE public.organization OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 50493)
-- Name: organization_id_organization_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.organization_id_organization_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_id_organization_seq OWNER TO postgres;

--
-- TOC entry 3174 (class 0 OID 0)
-- Dependencies: 200
-- Name: organization_id_organization_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.organization_id_organization_seq OWNED BY public.organization.id_organization;


--
-- TOC entry 211 (class 1259 OID 50636)
-- Name: post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post (
    id_post integer NOT NULL,
    name character varying(250) NOT NULL,
    salary numeric(36,2) NOT NULL,
    department_id integer NOT NULL
);


ALTER TABLE public.post OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 50634)
-- Name: post_id_post_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_id_post_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_post_seq OWNER TO postgres;

--
-- TOC entry 3175 (class 0 OID 0)
-- Dependencies: 210
-- Name: post_id_post_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_id_post_seq OWNED BY public.post.id_post;


--
-- TOC entry 221 (class 1259 OID 50829)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id_roles integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 50827)
-- Name: roles_id_roles_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_roles_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_roles_seq OWNER TO postgres;

--
-- TOC entry 3176 (class 0 OID 0)
-- Dependencies: 220
-- Name: roles_id_roles_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_roles_seq OWNED BY public.roles.id_roles;


--
-- TOC entry 213 (class 1259 OID 50650)
-- Name: strategy; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.strategy (
    id_strategy integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(250) NOT NULL,
    date_start date NOT NULL,
    date_end date NOT NULL,
    done boolean DEFAULT false,
    organization_id integer NOT NULL,
    date_done timestamp without time zone,
    log_del boolean DEFAULT false
);


ALTER TABLE public.strategy OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 50648)
-- Name: strategy_id_strategy_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.strategy_id_strategy_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.strategy_id_strategy_seq OWNER TO postgres;

--
-- TOC entry 3177 (class 0 OID 0)
-- Dependencies: 212
-- Name: strategy_id_strategy_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.strategy_id_strategy_seq OWNED BY public.strategy.id_strategy;


--
-- TOC entry 219 (class 1259 OID 50705)
-- Name: task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task (
    id_task integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(250) NOT NULL,
    date_start date NOT NULL,
    date_end date NOT NULL,
    done boolean DEFAULT false,
    employee_id integer NOT NULL,
    date_done timestamp without time zone,
    log_del boolean DEFAULT false NOT NULL
);


ALTER TABLE public.task OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 50703)
-- Name: task_id_task_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_id_task_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_task_seq OWNER TO postgres;

--
-- TOC entry 3178 (class 0 OID 0)
-- Dependencies: 218
-- Name: task_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_task_seq OWNED BY public.task.id_task;


--
-- TOC entry 2936 (class 2604 OID 50524)
-- Name: User id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User" ALTER COLUMN id_user SET DEFAULT nextval('public."User_id_user_seq"'::regclass);


--
-- TOC entry 2935 (class 2604 OID 50510)
-- Name: department id_department; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.department ALTER COLUMN id_department SET DEFAULT nextval('public.department_id_department_seq'::regclass);


--
-- TOC entry 2946 (class 2604 OID 50668)
-- Name: employee id_employee; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee ALTER COLUMN id_employee SET DEFAULT nextval('public.employee_id_employee_seq'::regclass);


--
-- TOC entry 2949 (class 2604 OID 50688)
-- Name: employee_post id_employee_post; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee_post ALTER COLUMN id_employee_post SET DEFAULT nextval('public.employee_post_id_employee_post_seq'::regclass);


--
-- TOC entry 2938 (class 2604 OID 50610)
-- Name: finances_operations id_operations; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finances_operations ALTER COLUMN id_operations SET DEFAULT nextval('public.finances_operations_id_operations_seq'::regclass);


--
-- TOC entry 2939 (class 2604 OID 50624)
-- Name: goal id_goal; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goal ALTER COLUMN id_goal SET DEFAULT nextval('public.goal_id_goal_seq'::regclass);


--
-- TOC entry 2933 (class 2604 OID 50498)
-- Name: organization id_organization; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization ALTER COLUMN id_organization SET DEFAULT nextval('public.organization_id_organization_seq'::regclass);


--
-- TOC entry 2942 (class 2604 OID 50639)
-- Name: post id_post; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post ALTER COLUMN id_post SET DEFAULT nextval('public.post_id_post_seq'::regclass);


--
-- TOC entry 2953 (class 2604 OID 50832)
-- Name: roles id_roles; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id_roles SET DEFAULT nextval('public.roles_id_roles_seq'::regclass);


--
-- TOC entry 2943 (class 2604 OID 50653)
-- Name: strategy id_strategy; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.strategy ALTER COLUMN id_strategy SET DEFAULT nextval('public.strategy_id_strategy_seq'::regclass);


--
-- TOC entry 2950 (class 2604 OID 50708)
-- Name: task id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id_task SET DEFAULT nextval('public.task_id_task_seq'::regclass);


--
-- TOC entry 3146 (class 0 OID 50521)
-- Dependencies: 205
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."User" (id_user, login, password, employee_id, roles_id) FROM stdin;
2	test	646667686a6b776567686b6a68666761484a303938424840bd001563085fc35165329ea1ff5c5ecbdbbeef	60	1
3	test2	646667686a6b776567686b6a68666761484a303938424840bd001563085fc35165329ea1ff5c5ecbdbbeef	61	2
4	test3	646667686a6b776567686b6a68666761484a303938424840bd001563085fc35165329ea1ff5c5ecbdbbeef	62	1
\.


--
-- TOC entry 3144 (class 0 OID 50507)
-- Dependencies: 203
-- Data for Name: department; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.department (id_department, name, description, organization_id) FROM stdin;
3	test3	test3 test3	19
1	test1	test1 test31	19
4	test4	test4 test4	19
5	Руководство	Данный отдел предназначен для сотрудников работающие в главном отделе организации	19
2	test2	test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2test2 test2	19
6	test6	dwuwdejed	19
\.


--
-- TOC entry 3156 (class 0 OID 50665)
-- Dependencies: 215
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employee (id_employee, surname, name, secondname, date_birth, seriapasp, numberpasp, email, department_id, emailverified) FROM stdin;
58	Администраторwadas	Дiojasfuiseifk	Админwaldk;askd	275760-08-09	3981	0898	emailhjdejkds	4	t
60	артем	артем	артем	1999-02-01	2312	53534	artemcesar88@gmail.com	1	t
61	артем	артем	артем	1999-02-01	2312	53584	isip_a.r.taran@mpt.ru	1	t
62	артем	артем	артем	1999-02-01	2312	53589	organizationcontrollapp@gmail.com	1	t
56	Особа2	Особенный2	Особен2	1011-01-01	1010	101010	10010	2	f
\.


--
-- TOC entry 3158 (class 0 OID 50685)
-- Dependencies: 217
-- Data for Name: employee_post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employee_post (id_employee_post, post_id, employee_id) FROM stdin;
65	7	58
66	6	58
24	2	56
\.


--
-- TOC entry 3148 (class 0 OID 50607)
-- Dependencies: 207
-- Data for Name: finances_operations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finances_operations (id_operations, summ, date_operation, description, organization_id) FROM stdin;
3	912839.00	1111-11-11	hhhhhh	19
1	9122.00	1111-11-11	hhhhhh	19
4	912839.00	1111-11-11	hhhhhh	19
5	-129830.00	2839-01-09	jkawkld	19
8	-1298300.00	2839-06-09	jkawkld	19
9	1234567.00	2839-09-09	jkawkld	19
6	-129830.00	2839-01-09	jkawkld	19
11	429830.00	2839-01-09	jkawkld	19
10	-82000.00	2839-03-09	jkawkld	19
12	-91239.00	1111-11-11	hhhhhh	19
13	2500000.00	2022-06-09	jkawkld	19
14	-12983005.00	2839-06-09	jkawkld	19
15	1293005.00	2839-10-10	jkawkld	19
7	82200000.00	2839-03-09	jkawkld	19
16	-8421395.86	2023-03-27	Зарплата сотрудников за Март	19
18	-8421395.86	2023-03-27	Зарплата сотрудников за Март	19
17	1111111.00	2023-08-27	Бонусы	19
\.


--
-- TOC entry 3150 (class 0 OID 50621)
-- Dependencies: 209
-- Data for Name: goal; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goal (id_goal, name, description, date_start, date_end, done, department_id, date_done, log_del) FROM stdin;
\.


--
-- TOC entry 3142 (class 0 OID 50495)
-- Dependencies: 201
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organization (id_organization, name, addres, inn, budget, auth_key, date_foundation) FROM stdin;
15	ТЕСТ ООО	лфдволыдОЛВЫДФОДЛВФОЦывф	219849124	9021.12	$2a$10$RIMOq4rr.DXxUdMkPqSrruZknlDqpW2JzV2J/7SHjmb87HXMblXHW	2023-07-02
17	ТЕСТ ООО	лфдволыдОЛВЫДФОДЛВФОЦывф	219849123	9021.12	$2a$10$62CzzL5wYVHfbyL3sp3iq./xQXlHPOSKYecO.Fu1TnOr.pX/7S0ja	2023-07-02
18	ТЕСТ ООО	лфдволыдОЛВЫДФОДЛВФОЦывф	219849121	9021.12	$2a$10$v5NmdTKMg2HbXPVD6GCJOS1Qd53pEZZ1XvUwTWRiCfqrnyGx6vSq	2023-07-02
19	Тест	Пушкино	777777777	59046317.28	$2a$10$XzGOM6xmgnDBsoM7r7lJyOlEwtNSQ5X6izphqFO8AarO7wcddqCW	2023-03-03
20	test4k	akwjldjklasjd	8127389213	8798231.00	$2a$10$ccKQAMHYHdVeQFLQ5M4fu2h7SfmOPi2ezjBjLlO0V7tSmTBgL1m	2393-08-19
\.


--
-- TOC entry 3152 (class 0 OID 50636)
-- Dependencies: 211
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.post (id_post, name, salary, department_id) FROM stdin;
5	Диктор	3921.31	4
6	Секретарь	3921.31	4
8	akjsdkjlwd	1230123908123.00	3
4	Директор	3921.31	5
3	Администратор	8321983.21	2
2	Уборщик	28139.12	3
7	Менеджер	3921.31	1
1	Менеджер	3921.31	3
\.


--
-- TOC entry 3162 (class 0 OID 50829)
-- Dependencies: 221
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id_roles, name) FROM stdin;
1	Сотрудник
2	Начальник отдела
3	Директор
\.


--
-- TOC entry 3154 (class 0 OID 50650)
-- Dependencies: 213
-- Data for Name: strategy; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.strategy (id_strategy, name, description, date_start, date_end, done, organization_id, date_done, log_del) FROM stdin;
38	testttttt	desc	1233-01-20	2999-01-21	f	19	2999-01-21 00:00:00	f
39	testttttt	desc	1233-01-20	2999-01-21	f	19	2999-01-21 11:11:11	f
\.


--
-- TOC entry 3160 (class 0 OID 50705)
-- Dependencies: 219
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.task (id_task, name, description, date_start, date_end, done, employee_id, date_done, log_del) FROM stdin;
\.


--
-- TOC entry 3179 (class 0 OID 0)
-- Dependencies: 204
-- Name: User_id_user_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."User_id_user_seq"', 4, true);


--
-- TOC entry 3180 (class 0 OID 0)
-- Dependencies: 202
-- Name: department_id_department_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.department_id_department_seq', 6, true);


--
-- TOC entry 3181 (class 0 OID 0)
-- Dependencies: 214
-- Name: employee_id_employee_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employee_id_employee_seq', 62, true);


--
-- TOC entry 3182 (class 0 OID 0)
-- Dependencies: 216
-- Name: employee_post_id_employee_post_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employee_post_id_employee_post_seq', 103, true);


--
-- TOC entry 3183 (class 0 OID 0)
-- Dependencies: 206
-- Name: finances_operations_id_operations_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finances_operations_id_operations_seq', 18, true);


--
-- TOC entry 3184 (class 0 OID 0)
-- Dependencies: 208
-- Name: goal_id_goal_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goal_id_goal_seq', 6, true);


--
-- TOC entry 3185 (class 0 OID 0)
-- Dependencies: 200
-- Name: organization_id_organization_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.organization_id_organization_seq', 20, true);


--
-- TOC entry 3186 (class 0 OID 0)
-- Dependencies: 210
-- Name: post_id_post_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_id_post_seq', 9, true);


--
-- TOC entry 3187 (class 0 OID 0)
-- Dependencies: 220
-- Name: roles_id_roles_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_roles_seq', 3, true);


--
-- TOC entry 3188 (class 0 OID 0)
-- Dependencies: 212
-- Name: strategy_id_strategy_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.strategy_id_strategy_seq', 39, true);


--
-- TOC entry 3189 (class 0 OID 0)
-- Dependencies: 218
-- Name: task_id_task_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.task_id_task_seq', 14, true);


--
-- TOC entry 2999 (class 2606 OID 50834)
-- Name: roles PK_11; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT "PK_11" PRIMARY KEY (id_roles);


--
-- TOC entry 2964 (class 2606 OID 50817)
-- Name: User UQ_1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "UQ_1" UNIQUE (login);


--
-- TOC entry 2966 (class 2606 OID 50819)
-- Name: User UQ_2; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "UQ_2" UNIQUE (employee_id);


--
-- TOC entry 2984 (class 2606 OID 50823)
-- Name: employee UQ_3; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT "UQ_3" UNIQUE (seriapasp, numberpasp);


--
-- TOC entry 2986 (class 2606 OID 50825)
-- Name: employee UQ_4; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT "UQ_4" UNIQUE (email);


--
-- TOC entry 2955 (class 2606 OID 50504)
-- Name: organization pk_1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT pk_1 PRIMARY KEY (id_organization);


--
-- TOC entry 2970 (class 2606 OID 50526)
-- Name: User pk_10; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT pk_10 PRIMARY KEY (id_user);


--
-- TOC entry 2989 (class 2606 OID 50670)
-- Name: employee pk_11; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT pk_11 PRIMARY KEY (id_employee);


--
-- TOC entry 2962 (class 2606 OID 50512)
-- Name: department pk_2; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT pk_2 PRIMARY KEY (id_department);


--
-- TOC entry 2994 (class 2606 OID 50690)
-- Name: employee_post pk_4; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee_post
    ADD CONSTRAINT pk_4 PRIMARY KEY (id_employee_post);


--
-- TOC entry 2973 (class 2606 OID 50612)
-- Name: finances_operations pk_5; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finances_operations
    ADD CONSTRAINT pk_5 PRIMARY KEY (id_operations);


--
-- TOC entry 2976 (class 2606 OID 50627)
-- Name: goal pk_6; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goal
    ADD CONSTRAINT pk_6 PRIMARY KEY (id_goal);


--
-- TOC entry 2979 (class 2606 OID 50641)
-- Name: post pk_7; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT pk_7 PRIMARY KEY (id_post);


--
-- TOC entry 2982 (class 2606 OID 50656)
-- Name: strategy pk_8; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.strategy
    ADD CONSTRAINT pk_8 PRIMARY KEY (id_strategy);


--
-- TOC entry 2997 (class 2606 OID 50711)
-- Name: task pk_9; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT pk_9 PRIMARY KEY (id_task);


--
-- TOC entry 2957 (class 2606 OID 50741)
-- Name: organization unique_1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT unique_1 UNIQUE (inn);


--
-- TOC entry 2959 (class 2606 OID 50743)
-- Name: organization unique_2; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT unique_2 UNIQUE (auth_key);


--
-- TOC entry 2960 (class 1259 OID 50518)
-- Name: fk_21; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_21 ON public.department USING btree (organization_id);


--
-- TOC entry 2987 (class 1259 OID 50682)
-- Name: fk_25; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_25 ON public.employee USING btree (department_id);


--
-- TOC entry 2967 (class 1259 OID 50753)
-- Name: fk_27; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_27 ON public."User" USING btree (employee_id);


--
-- TOC entry 2995 (class 1259 OID 50717)
-- Name: fk_3; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_3 ON public.task USING btree (employee_id);


--
-- TOC entry 2980 (class 1259 OID 50662)
-- Name: fk_4; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_4 ON public.strategy USING btree (organization_id);


--
-- TOC entry 2968 (class 1259 OID 50840)
-- Name: fk_48; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_48 ON public."User" USING btree (roles_id);


--
-- TOC entry 2977 (class 1259 OID 50647)
-- Name: fk_5; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_5 ON public.post USING btree (department_id);


--
-- TOC entry 2974 (class 1259 OID 50633)
-- Name: fk_6; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_6 ON public.goal USING btree (department_id);


--
-- TOC entry 2990 (class 1259 OID 50760)
-- Name: fk_66; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_66 ON public.employee_post USING btree (post_id);


--
-- TOC entry 2971 (class 1259 OID 50618)
-- Name: fk_7; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_7 ON public.finances_operations USING btree (organization_id);


--
-- TOC entry 2991 (class 1259 OID 50702)
-- Name: fk_8; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_8 ON public.employee_post USING btree (employee_id);


--
-- TOC entry 2992 (class 1259 OID 50701)
-- Name: fk_9; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_9 ON public.employee_post USING btree (post_id);


--
-- TOC entry 3004 (class 2606 OID 50776)
-- Name: goal department_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goal
    ADD CONSTRAINT department_id FOREIGN KEY (department_id) REFERENCES public.department(id_department) ON DELETE CASCADE;


--
-- TOC entry 3005 (class 2606 OID 50781)
-- Name: post department_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT department_id FOREIGN KEY (department_id) REFERENCES public.department(id_department) ON DELETE CASCADE;


--
-- TOC entry 3007 (class 2606 OID 50786)
-- Name: employee department_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT department_id FOREIGN KEY (department_id) REFERENCES public.department(id_department) ON DELETE CASCADE;


--
-- TOC entry 3010 (class 2606 OID 50791)
-- Name: task employee_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT employee_id FOREIGN KEY (employee_id) REFERENCES public.employee(id_employee) ON DELETE CASCADE;


--
-- TOC entry 3001 (class 2606 OID 50796)
-- Name: User employee_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT employee_id FOREIGN KEY (employee_id) REFERENCES public.employee(id_employee) ON DELETE CASCADE;


--
-- TOC entry 3008 (class 2606 OID 50801)
-- Name: employee_post employee_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee_post
    ADD CONSTRAINT employee_id FOREIGN KEY (employee_id) REFERENCES public.employee(id_employee) ON DELETE CASCADE;


--
-- TOC entry 3002 (class 2606 OID 50835)
-- Name: User fk_44; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT fk_44 FOREIGN KEY (roles_id) REFERENCES public.roles(id_roles);


--
-- TOC entry 3006 (class 2606 OID 50761)
-- Name: strategy organization_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.strategy
    ADD CONSTRAINT organization_id FOREIGN KEY (organization_id) REFERENCES public.organization(id_organization) ON DELETE CASCADE;


--
-- TOC entry 3003 (class 2606 OID 50766)
-- Name: finances_operations organization_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finances_operations
    ADD CONSTRAINT organization_id FOREIGN KEY (organization_id) REFERENCES public.organization(id_organization) ON DELETE CASCADE;


--
-- TOC entry 3000 (class 2606 OID 50771)
-- Name: department organization_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT organization_id FOREIGN KEY (organization_id) REFERENCES public.organization(id_organization) ON DELETE CASCADE;


--
-- TOC entry 3009 (class 2606 OID 50806)
-- Name: employee_post post_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee_post
    ADD CONSTRAINT post_id FOREIGN KEY (post_id) REFERENCES public.post(id_post) ON DELETE CASCADE;


-- Completed on 2023-05-10 14:28:54

--
-- PostgreSQL database dump complete
--

