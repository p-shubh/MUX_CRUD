--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 9.6.5

-- Started on 2017-11-24 22:20:37

SET
statement_timeout = 0;
SET
lock_timeout = 0;
SET
idle_in_transaction_session_timeout = 0;
SET
client_encoding = 'UTF8';
SET
standard_conforming_strings = on;
SET
check_function_bodies = false;
SET
client_min_messages = warning;
SET
row_security = off;

SET
search_path = public, pg_catalog;

--
-- TOC entry 185 (class 1259 OID 16429)
-- Name: account_metrics_consecutive_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE account_metrics_consecutive_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE CACHE 1;


ALTER TABLE account_metrics_consecutive_seq OWNER TO postgres;

SET
default_tablespace = '';

SET
default_with_oids = false;

--
-- TOC entry 186 (class 1259 OID 16431)
-- Name: account_metrics; Type: TABLE; Schema: public; Owner: postgres
--
/* bpr table indicates to calculate gap and penetration for the BPR for different location */
CREATE TABLE bpr
(
    date              TEXT,
    department        TEXT,
    part_number       TEXT, /* Unique assembly id */
    norm              float,
    stock             float,
    stock_2           float,
    gap               float, /* BPRgap = norm - stock */
    penetration       float, /* penetration = gap/norm  *100 */
    Colour            TEXT,
    Colour_2          TEXT,
    Total_Coverage    float,
    Desired_inventory float,
    Actual_inventory  float,
    White_inventory   float,
    Created_By        int,
    Created_At        text,
    Updated_By        int,
    Updated_At        text
);
/* norm table indicates norm input for parts for different location */
CREATE TABLE norm
(
    date        TEXT,
    department  TEXT,
    part_number TEXT,
    item_class  TEXT,
    amr         float,
    norm        float,
    Created_By  int,
    Created_At  text,
    Updated_By  int,
    Updated_At  text
);
/* Part table indicates top level assembly structure . */
CREATE TABLE part
(
    Part_ID     SERIAL PRIMARY KEY,
    part_number TEXT, /* Unique assembly id */
    description TEXT, /* assembly details. */
    Created_By  int,
    Created_At  text,
    Updated_By  int,
    Updated_At  text,
    Is_deleted  int

);

CREATE TABLE part_details
(
    Part_Details_ID       SERIAL PRIMARY KEY,
    Part_ID               int,
    part_number           TEXT,
    department            TEXT,
    sub_category          TEXT,
    Item_Class            TEXT,
    rate                  float,
    Alternate_part_number TEXT,
    Plant                 TEXT,
    Created_By            int,
    Created_At            text,
    Updated_By            int,
    Updated_At            text,
    Row_Order             int,
    Is_deleted            int
);

/* shift table indicates work schedule between start & end time */
CREATE TABLE shift
(
    start_time TEXT,
    end_time   TEXT,
    index      int,
    CONSTRAINT shift_key PRIMARY KEY (index)
);

CREATE TABLE AMR
(
    date               TEXT,
    department         TEXT,
    File_Type          TEXT,
    part_number        TEXT,
    part_description   TEXT,
    To_Date_Aval       float,
    Net_Month_Shortage float,
    AMR                float,
    Created_By         int,
    Created_At         text,
    Updated_By         int,
    Updated_At         text
);

CREATE TABLE Norm_CD
(
    Norm_CD_ID    SERIAL PRIMARY KEY,
    Part_ID       int,
    Department_ID int,
    date          TEXT,
    department    TEXT,
    part_number   TEXT,
    NOD           float,
    lead_time     float,
    safety_factor float,
    Buyer         TEXT,
    Supplier      TEXT,
    Created_By    int,
    Created_At    text,
    Updated_By    int,
    Updated_At    text,
    Row_Order     int,
    Is_deleted    int
);

CREATE TABLE Stock
(
    date                 TEXT,
    department           TEXT,
    part_number          TEXT,
    todate_aval          float,
    under_inspection     float,
    store_stock          float,
    vmi_stock            float,
    sub_contractor_stock float,
    total_stock          float,
    total_stock_2        float,
    Created_By           int,
    Created_At           text,
    Updated_By           int,
    Updated_At           text

);

CREATE TABLE Stock_file_log
(
    date       TEXT,
    department TEXT,
    file_type  TEXT,
    file_name  TEXT,
    Created_By int,
    Created_At text,
    Updated_By int,
    Updated_At text

);

CREATE TABLE Stock_file_Map
(

    department      TEXT,
    stock_file_type TEXT,
    Created_By      int,
    Created_At      text,
    Updated_By      int,
    Updated_At      text

);

CREATE TABLE Full_kit_master
(
    department              TEXT,
    FG_assembly_code        TEXT,
    FG_assembly_description TEXT,
    Child_Part_Code         TEXT,
    Line                    TEXT,
    Created_By              int,
    Created_At              text,
    Updated_By              int,
    Updated_At              text
);

CREATE TABLE Full_kit_output
(
    date             TEXT,
    department       TEXT,
    FG_assembly_code TEXT,
    Remark           int,
    Item_Class       TEXT,
    Created_By       int,
    Created_At       text,
    Updated_By       int,
    Updated_At       text
);

CREATE TABLE Part_supplier_details
(
    Part_supplier_details_ID SERIAL PRIMARY KEY,
    Part_ID                  int,
    Department_ID            int,
    Sub_category_ID          int,
    part_number              TEXT,
    department               TEXT,
    sub_category             TEXT,
    buyer                    TEXT,
    supplier                 TEXT,
    display_name             TEXT,
    Created_By               int,
    Created_At               text,
    Updated_By               int,
    Updated_At               text,
    Row_Order                int,
    Is_deleted               int
);

CREATE TABLE import_norm
(
    date        TEXT,
    department  TEXT,
    part_number TEXT,
    item_class  TEXT,
    AMR_1       float,
    AMR_2       float,
    AMR_3       float,
    norm_1      float,
    norm_2      float,
    norm_3      float,
    Created_By  int,
    Created_At  text,
    Updated_By  int,
    Updated_At  text
);

CREATE TABLE import_AMR
(
    date             TEXT,
    department       TEXT,
    part_number      TEXT,
    part_description TEXT,
    AMR_1            float,
    AMR_2            float,
    AMR_3            float,
    Created_By       int,
    Created_At       text,
    Updated_By       int,
    Updated_At       text
);

CREATE TABLE Department
(
    DepartmentID  SERIAL PRIMARY KEY,
    department    TEXT,
    display_order int,
    Created_By    int,
    Created_At    text,
    Updated_By    int,
    Updated_At    text
);

CREATE TABLE Sub_Category
(
    Sub_CategoryID SERIAL PRIMARY KEY,
    department     TEXT,
    sub_category   TEXT,
    display_order  int,
    DepartmentID   int,
    Created_By     int,
    Created_At     text,
    Updated_By     int,
    Updated_At     text
);

CREATE TABLE Item_class
(
    Item_classID   SERIAL PRIMARY KEY,
    Sub_CategoryID int,
    DepartmentID   int,
    department     TEXT,
    sub_category   TEXT,
    item_class     TEXT,
    Created_By     int,
    Created_At     text,
    Updated_By     int,
    Updated_At     text
);

CREATE TABLE Part_supplier_output
(
    department     TEXT,
    buyer          TEXT,
    supplier       TEXT,
    count_of_red   int,
    count_of_black int,
    Created_By     int,
    Created_At     text,
    Updated_By     int,
    Updated_At     text
);

CREATE TABLE H_Modules
(
    ID               SERIAL PRIMARY KEY,
    Name             TEXT,
    DisplayIndex     int,
    DefaultPageIndex int,
    ModuleIcon       TEXT
);

CREATE TABLE H_Pages
(
    ID                 SERIAL PRIMARY KEY,
    Name               TEXT,
    NameOnMenu         TEXT,
    Description        TEXT,
    DisplayIndex       int,
    DefaultModuleID    int,
    DefaultSubModuleID int,
    ActualPagePath     TEXT,
    PageType           TEXT,
    PageID             int
);

CREATE TABLE H_SubModules
(
    ID            SERIAL PRIMARY KEY,
    Name          TEXT,
    ModuleID      int,
    DisplayIndex  int,
    DefaultPageID int
);

CREATE TABLE H_SubPages
(
    ID     SERIAL PRIMARY KEY,
    PageID int,
    Name   TEXT
);

CREATE TABLE M_Employees
(
    ID                    SERIAL PRIMARY KEY,
    Code                  TEXT,
    Name                  TEXT,
    MiddleName            TEXT,
    LastName              TEXT,
    EmployeeGender        TEXT,
    GSTin                 TEXT,
    DateOfBirth           TEXT,
    JoiningDate           TEXT,
    DesignationID         int,
    Address               TEXT,
    ActualAttendanceHours decimal,
    City                  TEXT,
    State                 TEXT,
    PIN                   TEXT,
    Phone                 TEXT,
    Email                 TEXT,
    EmployeeType          int,
    IsDriver              bit,
    IsNonEmployee         bit,
    IsActive              bit,
    PartiesID             int,
    CompanyID             int,
    DivisionID            int,
    CreatedBy             int,
    CreatedOn             TIME,
    UpdatedBy             int,
    UpdatedOn             TIME
);

CREATE TABLE M_ManagementUsers
(
    ID          SERIAL PRIMARY KEY,
    Name        TEXT,
    MiddleName  TEXT,
    LastName    TEXT,
    Address     TEXT,
    City        TEXT,
    State       TEXT,
    PIN         TEXT,
    GSTin       TEXT,
    Phone       TEXT,
    Email       TEXT,
    PartiesID   int,
    CompanyID   int,
    IsEmailSent int,
    CreatedBy   int,
    CreatedOn   TIME,
    UpdatedBy   int,
    UpdatedOn   TIME
);

CREATE TABLE M_RoleAccess
(
    ID            SERIAL PRIMARY KEY,
    RoleID        int,
    CompanyID     int,
    DivisionID    int,
    ModuleID      int,
    SubModuleID   int,
    PageID        int,
    IsShow        int,
    IsShowAddPage int,
    IsAdd         int,
    IsEditSelf    int,
    IsEdit        int,
    IsView        int,
    IsDeleteSelf  int,
    IsDelete      int,
    IsPrint       int
);

CREATE TABLE M_Roles
(
    ID          SERIAL PRIMARY KEY,
    Name        TEXT,
    Description TEXT,
    IsActive    TEXT,
    Page_Path   TEXT
);

CREATE TABLE M_UserAccess
(
    ID            SERIAL PRIMARY KEY,
    UserID        int,
    CompanyID     int,
    DivisionID    int,
    ModuleID      int,
    SubModuleID   int,
    PageID        int,
    IsShow        int,
    IsShowAddPage int,
    IsAdd         int,
    IsEditSelf    int,
    IsEdit        int,
    IsView        int,
    IsDeleteSelf  int,
    IsDelete      int,
    IsPrint       int
);

CREATE TABLE M_Users
(
    UserID          SERIAL PRIMARY KEY,
    FirstName       TEXT,
    LastName        TEXT,
    LoginName       TEXT,
    Mobile          TEXT,
    EmailID         TEXT,
    Password        TEXT,
    IsResetPassword int,
    CustomerID      int,
    CreatedBy       int,
    CreatedOn       TEXT,
    UpdatedBy       int,
    UpdatedOn       TEXT
);

CREATE TABLE MC_UserDivisions
(
    ID           SERIAL PRIMARY KEY,
    UserID       int,
    DepartmentID int,
    Row_Order    int
);

CREATE TABLE MC_UserRoles
(
    ID        SERIAL PRIMARY KEY,
    UserID    int,
    RoleID    int,
    Row_Order int
);

CREATE TABLE C_Companies
(
    ID                   SERIAL PRIMARY KEY,
    Name                 TEXT,
    CompanyType          TEXT,
    GSTIN                TEXT,
    FSSAI                TEXT,
    Address              TEXT,
    Email                TEXT,
    PAN                  TEXT,
    StateID              int,
    CustomerCareNo       TEXT,
    BankName             TEXT,
    IFSC                 TEXT,
    AccountNo            TEXT,
    CreatedBy            int,
    CreatedOn            timestamp,
    UpdatedBy            int,
    UpdatedOn            timestamp,
    DivisionAbbreviation TEXT,
    PinCode              TEXT,
    LogoName             TEXT,
    EwayMailID           TEXT
);

CREATE TABLE Plant
(
    ID        SERIAL PRIMARY KEY,
    PlantName text
);

CREATE TABLE Buyer
(
    ID        SERIAL PRIMARY KEY,
    BuyerName text
);

CREATE TABLE Supplier
(
    ID           SERIAL PRIMARY KEY,
    SupplierName text
);

CREATE TABLE File_Format
(
    Department      TEXT,
    FunctionName    TEXT,
    FileName        TEXT,
    Excel_File_Name TEXT,
    ColumnName      TEXT,
    Column_Order    int,
    Column_Number   int
);

CREATE TABLE Line
(
    LineID     SERIAL PRIMARY KEY,
    LineName   TEXT,
    Department TEXT
);



-- Completed on 2017-11-24 22:20:39

--
-- PostgreSQL database dump complete
--
