CREATE TABLE fleetguard_trigger
(
    date                  TEXT,
    location              TEXT,
    organization          TEXT,
    item_code             TEXT,
    description           TEXT,
    item_location         TEXT,
    drg_rev               TEXT,
    priority              TEXT,
    bp                    int,
    po_no                 TEXT,
    release_no            TEXT,
    release_method        TEXT,
    release_date          TEXT,
    open_pending_quantity int,
    asn_qty               TEXT,
    CONSTRAINT trigger_key PRIMARY KEY (date, item_code, location)
);
/* production_data table indicates daily production report */
CREATE TABLE production_data
(
    operator       TEXT,
    date           TEXT,
    shift          TEXT,
    timeslot       TEXT,
    partNumber     TEXT,
    subPartNumber  TEXT,
    operation      TEXT,
    machine        TEXT,
    planned        int,
    produced       int,
    rejected       int,
    downtime       int,
    downtimeReason TEXT,
    remarks        TEXT
);

/* stock_RM table indicates amount of closing_stock (unsold stock) at the end of reporting period */ 
CREATE TABLE stock_RawMaterial
(
    date          TEXT,
    location      TEXT,
    grade         TEXT,
    thickness     float,
    width         int,
    length        int,
    closing_stock int
);
/* stock_FG table indicates amount of closing_stock input for parts for different location */
CREATE TABLE stock_FG
(
    date          TEXT,
    location      TEXT,
    part_number   TEXT,
    closing_stock int
);
/* stock_WIP table indicates amount of closing_stock input for parts for different location */
CREATE TABLE stock_WIP
(
    date          TEXT,
    location      TEXT,
    part_number   TEXT,
    closing_stock int
);
/* bpr table indicates to calculate gap and penetration for the BPR for different location */
CREATE TABLE bpr
(
    date          TEXT,
    department      TEXT,
    part_number   TEXT,     /* Unique assembly id */
    gap           int,      /* BPRgap = norm - stock */
    penetration   int       /* penetration = gap/norm  *100 */
);
/* norm table indicates norm input for parts for different location */
CREATE TABLE norm
(
    date           TEXT,
    department     TEXT,
    part_number    TEXT,
    norm           int

);
/* Part table indicates top level assembly structure . */
CREATE TABLE part (
	part_number TEXT,   /* Unique assembly id */
	description TEXT,   /* assembly details. */
    CONSTRAINT part_key PRIMARY KEY (part_number)
);


--  Assumption - Subassembly number is not dependent on location.

/*  Subassembly is level 2 table for a part hierearchy */ 
CREATE TABLE subassembly (
        sa_number TEXT,    /* subasssembly/child part for the assembly */
        assembly  TEXT,    /* part or assembly number */
        description TEXT,  /* description of the subassembly */
        CONSTRAINT sa_key PRIMARY KEY (sa_number)
);
/* shift table indicates work schedule between start & end time */
CREATE TABLE shift (
	start_time TEXT,
	end_time TEXT,
	index int,
    CONSTRAINT shift_key PRIMARY KEY (index)
);

CREATE TABLE AMR (
    date          TEXT,
   part_number   TEXT,
    AMR           int
    
);

CREATE TABLE Stock 
(
   date                 TEXT,
   department           TEXT,
   part_number          TEXT,
   under_inspection     int,
   store_stock          int,
   vmi_stock            int,
   sub_contractor_stock int,
   total_stock          int

);

