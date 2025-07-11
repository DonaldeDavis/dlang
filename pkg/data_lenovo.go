// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMaplenovo = map[string]*DeviceData{
    "Flex System": {
        Manufacturer: "Lenovo",
        Model: "Flex System",
        Slug: "lenovo-flex-system",
        UHeight: 10,
        PartNumber: "8721W3X",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU3", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU4", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU5", Label: "", Type: "iec-60320-c20", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
            { Name: "Bay 1", Label: "" },
            { Name: "Bay 2", Label: "" },
            { Name: "Bay 3", Label: "" },
            { Name: "Bay 4", Label: "" },
            { Name: "Bay 5", Label: "" },
            { Name: "Bay 6", Label: "" },
            { Name: "Bay 7", Label: "" },
            { Name: "Bay 8", Label: "" },
            { Name: "Bay 9", Label: "" },
            { Name: "Bay 10", Label: "" },
            { Name: "Bay 11", Label: "" },
            { Name: "Bay 12", Label: "" },
            { Name: "Bay 13", Label: "" },
            { Name: "Bay 14", Label: "" },
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "CMM", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "RackSwitch G8052": {
        Manufacturer: "Lenovo",
        Model: "RackSwitch G8052",
        Slug: "lenovo-rackswitch-g8052",
        UHeight: 1,
        PartNumber: "715952F",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: true,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 23.1,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Serial", Type: "usb-mini-a", Label: "Serial", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "FAN1", Label: "FAN1", Position: "FAN1" },
            { Name: "FAN2", Label: "FAN2", Position: "FAN2" },
            { Name: "FAN3", Label: "FAN3", Position: "FAN3" },
            { Name: "FAN4", Label: "FAN4", Position: "FAN4" },
            { Name: "PWR1", Label: "PWR1", Position: "PWR1" },
            { Name: "PWR2", Label: "PWR2", Position: "PWR2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "1", Label: "1", Type: "1000base-tx", MgmtOnly: false },
            { Name: "2", Label: "2", Type: "1000base-tx", MgmtOnly: false },
            { Name: "3", Label: "3", Type: "1000base-tx", MgmtOnly: false },
            { Name: "4", Label: "4", Type: "1000base-tx", MgmtOnly: false },
            { Name: "5", Label: "5", Type: "1000base-tx", MgmtOnly: false },
            { Name: "6", Label: "6", Type: "1000base-tx", MgmtOnly: false },
            { Name: "7", Label: "7", Type: "1000base-tx", MgmtOnly: false },
            { Name: "8", Label: "8", Type: "1000base-tx", MgmtOnly: false },
            { Name: "9", Label: "9", Type: "1000base-tx", MgmtOnly: false },
            { Name: "10", Label: "10", Type: "1000base-tx", MgmtOnly: false },
            { Name: "11", Label: "11", Type: "1000base-tx", MgmtOnly: false },
            { Name: "12", Label: "12", Type: "1000base-tx", MgmtOnly: false },
            { Name: "13", Label: "13", Type: "1000base-tx", MgmtOnly: false },
            { Name: "14", Label: "14", Type: "1000base-tx", MgmtOnly: false },
            { Name: "15", Label: "15", Type: "1000base-tx", MgmtOnly: false },
            { Name: "16", Label: "16", Type: "1000base-tx", MgmtOnly: false },
            { Name: "17", Label: "17", Type: "1000base-tx", MgmtOnly: false },
            { Name: "18", Label: "18", Type: "1000base-tx", MgmtOnly: false },
            { Name: "19", Label: "19", Type: "1000base-tx", MgmtOnly: false },
            { Name: "20", Label: "20", Type: "1000base-tx", MgmtOnly: false },
            { Name: "21", Label: "21", Type: "1000base-tx", MgmtOnly: false },
            { Name: "22", Label: "22", Type: "1000base-tx", MgmtOnly: false },
            { Name: "23", Label: "23", Type: "1000base-tx", MgmtOnly: false },
            { Name: "24", Label: "24", Type: "1000base-tx", MgmtOnly: false },
            { Name: "25", Label: "25", Type: "1000base-tx", MgmtOnly: false },
            { Name: "26", Label: "26", Type: "1000base-tx", MgmtOnly: false },
            { Name: "27", Label: "27", Type: "1000base-tx", MgmtOnly: false },
            { Name: "28", Label: "28", Type: "1000base-tx", MgmtOnly: false },
            { Name: "29", Label: "29", Type: "1000base-tx", MgmtOnly: false },
            { Name: "30", Label: "30", Type: "1000base-tx", MgmtOnly: false },
            { Name: "31", Label: "31", Type: "1000base-tx", MgmtOnly: false },
            { Name: "32", Label: "32", Type: "1000base-tx", MgmtOnly: false },
            { Name: "33", Label: "33", Type: "1000base-tx", MgmtOnly: false },
            { Name: "34", Label: "34", Type: "1000base-tx", MgmtOnly: false },
            { Name: "35", Label: "35", Type: "1000base-tx", MgmtOnly: false },
            { Name: "36", Label: "36", Type: "1000base-tx", MgmtOnly: false },
            { Name: "37", Label: "37", Type: "1000base-tx", MgmtOnly: false },
            { Name: "38", Label: "38", Type: "1000base-tx", MgmtOnly: false },
            { Name: "39", Label: "39", Type: "1000base-tx", MgmtOnly: false },
            { Name: "40", Label: "40", Type: "1000base-tx", MgmtOnly: false },
            { Name: "41", Label: "41", Type: "1000base-tx", MgmtOnly: false },
            { Name: "42", Label: "42", Type: "1000base-tx", MgmtOnly: false },
            { Name: "43", Label: "43", Type: "1000base-tx", MgmtOnly: false },
            { Name: "44", Label: "44", Type: "1000base-tx", MgmtOnly: false },
            { Name: "45", Label: "45", Type: "1000base-tx", MgmtOnly: false },
            { Name: "46", Label: "46", Type: "1000base-tx", MgmtOnly: false },
            { Name: "47", Label: "47", Type: "1000base-tx", MgmtOnly: false },
            { Name: "48", Label: "48", Type: "1000base-tx", MgmtOnly: false },
            { Name: "XGE1", Label: "49", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "XGE2", Label: "50", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "XGE3", Label: "51", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "XGE4", Label: "52", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
    "System x3250 M6": {
        Manufacturer: "Lenovo",
        Model: "System x3250 M6",
        Slug: "lenovo-system-x3250-m6",
        UHeight: 1,
        PartNumber: "3633",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU 2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "System x3550 M5": {
        Manufacturer: "Lenovo",
        Model: "System x3550 M5",
        Slug: "lenovo-system-x3550-m5",
        UHeight: 1,
        PartNumber: "5463AC1",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU 2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "IMM", Label: "", Type: "100base-tx", MgmtOnly: true },
        },
    },
    "System x3650 M5": {
        Manufacturer: "Lenovo",
        Model: "System x3650 M5",
        Slug: "lenovo-system-x3650-m5",
        UHeight: 2,
        PartNumber: "8871",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 19,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU 1", Label: "", Position: "1" },
            { Name: "PSU 2", Label: "", Position: "2" },
            { Name: "PCIe slot 1 - FH", Label: "1", Position: "PCIe-1" },
            { Name: "PCIe slot 2 - FH", Label: "2", Position: "PCIe-2" },
            { Name: "PCIe slot 3 or ML2 - FH", Label: "3", Position: "PCIe-3" },
            { Name: "PCIe slot 4 - LP", Label: "4", Position: "PCIe-4" },
            { Name: "PCIe slot 5 - LP", Label: "5", Position: "PCIe-5" },
            { Name: "PCIe slot 6 - FH", Label: "6", Position: "PCIe-6" },
            { Name: "PCIe slot 7 - FH", Label: "7", Position: "PCIe-7" },
            { Name: "PCIe slot 8 - FH", Label: "8", Position: "PCIe-8" },
            { Name: "Mezzanine LOM (ML2)", Label: "LOM", Position: "PCIe-8" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "IMM", Label: "", Type: "100base-tx", MgmtOnly: true },
        },
    },
    "System x3750 M4": {
        Manufacturer: "Lenovo",
        Model: "System x3750 M4",
        Slug: "lenovo-system-x3750-m4",
        UHeight: 2,
        PartNumber: "8722C1U",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU 2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "IMM", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkAgile MX3531-F CN": {
        Manufacturer: "Lenovo",
        Model: "ThinkAgile MX3531-F CN",
        Slug: "lenovo-thinkagile-mx3531-f-cn",
        UHeight: 2,
        PartNumber: "7D66",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 38.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU-1", Label: "", Position: "PSU-1" },
            { Name: "PSU-2", Label: "", Position: "PSU-2" },
            { Name: "OCP slot 1", Label: "", Position: "OCP-1" },
            { Name: "Riser 1 - PCIe slot 1", Label: "1 PCIe", Position: "PCIe-1" },
            { Name: "Riser 1 - PCIe slot 2", Label: "2 PCIe", Position: "PCIe-2" },
            { Name: "Riser 1 - PCIe slot 3", Label: "3 PCIe", Position: "PCIe-3" },
            { Name: "Riser 2 - PCIe slot 4", Label: "4 PCIe", Position: "PCIe-4" },
            { Name: "Riser 2 - PCIe slot 5", Label: "5 PCIe", Position: "PCIe-5" },
            { Name: "Riser 2 - PCIe slot 6", Label: "6 PCIe", Position: "PCIe-6" },
            { Name: "Riser 3 - PCIe slot 7", Label: "7 PCIe", Position: "PCIe-7" },
            { Name: "Riser 3 - PCIe slot 8", Label: "8 PCIe", Position: "PCIe-8" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkServer RD230": {
        Manufacturer: "Lenovo",
        Model: "ThinkServer RD230",
        Slug: "lenovo-thinkserver-rd230",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 14,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem 10Gb 2-port SFP&#43; LOM": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem 10Gb 2-port SFP&#43; LOM",
        Slug: "lenovo-thinksystem-10gb-2-port-sfpp-lom",
        UHeight: 0,
        PartNumber: "7ZT7A00546",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "1", Label: "1", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "2", Label: "2", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
    "ThinkSystem 10Gb 4-port SFP&#43; LOM": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem 10Gb 4-port SFP&#43; LOM",
        Slug: "lenovo-thinksystem-10gb-4-port-sfpp-lom",
        UHeight: 0,
        PartNumber: "7ZT7A00547",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "1", Label: "1", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "2", Label: "2", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "3", Label: "3", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "4", Label: "4", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
    "ThinkSystem SR250 V2": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR250 V2",
        Slug: "lenovo-thinksystem-sr250-v2",
        UHeight: 1,
        PartNumber: "7D7Q",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 12.3,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU1", Label: "", Position: "1" },
            { Name: "PSU2", Label: "", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "eno1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eno2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR530": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR530",
        Slug: "lenovo-thinksystem-sr530",
        UHeight: 1,
        PartNumber: "7X07",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR550": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR550",
        Slug: "lenovo-thinksystem-sr550",
        UHeight: 2,
        PartNumber: "7X04",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR630": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR630",
        Slug: "lenovo-thinksystem-sr630",
        UHeight: 1,
        PartNumber: "7X02",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR645": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR645",
        Slug: "lenovo-thinksystem-sr645",
        UHeight: 1,
        PartNumber: "7D2X",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 20.2,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU 1", Label: "", Position: "PSU1" },
            { Name: "PSU 2", Label: "", Position: "PSU2" },
            { Name: "OCP3", Label: "", Position: "OCP3" },
            { Name: "PCIe1", Label: "", Position: "Slot1" },
            { Name: "PCIe2", Label: "", Position: "Slot2" },
            { Name: "PCIe3", Label: "", Position: "Slot3" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR650": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR650",
        Slug: "lenovo-thinksystem-sr650",
        UHeight: 2,
        PartNumber: "7X06",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR650 V2": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR650 V2",
        Slug: "lenovo-thinksystem-sr650-v2",
        UHeight: 2,
        PartNumber: "7Z7x",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 38.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU-1", Label: "", Position: "PSU-1" },
            { Name: "PSU-2", Label: "", Position: "PSU-2" },
            { Name: "OCP slot 1", Label: "", Position: "OCP-1" },
            { Name: "Riser 1 - PCIe slot 1", Label: "1 PCIe", Position: "PCIe-1" },
            { Name: "Riser 1 - PCIe slot 2", Label: "2 PCIe", Position: "PCIe-2" },
            { Name: "Riser 1 - PCIe slot 3", Label: "3 PCIe", Position: "PCIe-3" },
            { Name: "Riser 2 - PCIe slot 4", Label: "4 PCIe", Position: "PCIe-4" },
            { Name: "Riser 2 - PCIe slot 5", Label: "5 PCIe", Position: "PCIe-5" },
            { Name: "Riser 2 - PCIe slot 6", Label: "6 PCIe", Position: "PCIe-6" },
            { Name: "Riser 3 - PCIe slot 7", Label: "7 PCIe", Position: "PCIe-7" },
            { Name: "Riser 3 - PCIe slot 8", Label: "8 PCIe", Position: "PCIe-8" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR655 V3": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR655 V3",
        Slug: "lenovo-thinksystem-sr655-v3",
        UHeight: 2,
        PartNumber: "7D9x",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 38.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU-1", Label: "", Position: "PSU-1" },
            { Name: "PSU-2", Label: "", Position: "PSU-2" },
            { Name: "OCP slot 1", Label: "", Position: "OCP-1" },
            { Name: "Riser 1 - PCIe slot 1", Label: "1 PCIe", Position: "PCIe-1" },
            { Name: "Riser 1 - PCIe slot 2", Label: "2 PCIe", Position: "PCIe-2" },
            { Name: "Riser 1 - PCIe slot 3", Label: "3 PCIe", Position: "PCIe-3" },
            { Name: "Riser 2 - PCIe slot 4", Label: "4 PCIe", Position: "PCIe-4" },
            { Name: "Riser 2 - PCIe slot 5", Label: "5 PCIe", Position: "PCIe-5" },
            { Name: "Riser 2 - PCIe slot 6", Label: "6 PCIe", Position: "PCIe-6" },
            { Name: "Riser 3 - PCIe slot 7", Label: "7 PCIe", Position: "PCIe-7" },
            { Name: "Riser 3 - PCIe slot 8", Label: "8 PCIe", Position: "PCIe-8" },
            { Name: "Riser 4 - PCIe slot 9", Label: "9 PCIe", Position: "PCIe-9" },
            { Name: "Riser 4 - PCIe slot 10", Label: "10 PCIe", Position: "PCIe-10" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR665 V3": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR665 V3",
        Slug: "lenovo-thinksystem-sr665-v3",
        UHeight: 2,
        PartNumber: "7D9x",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 38.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU-1", Label: "", Position: "PSU-1" },
            { Name: "PSU-2", Label: "", Position: "PSU-2" },
            { Name: "OCP slot 1", Label: "", Position: "OCP-1" },
            { Name: "Riser 1 - PCIe slot 1", Label: "1 PCIe", Position: "PCIe-1" },
            { Name: "Riser 1 - PCIe slot 2", Label: "2 PCIe", Position: "PCIe-2" },
            { Name: "Riser 1 - PCIe slot 3", Label: "3 PCIe", Position: "PCIe-3" },
            { Name: "Riser 2 - PCIe slot 4", Label: "4 PCIe", Position: "PCIe-4" },
            { Name: "Riser 2 - PCIe slot 5", Label: "5 PCIe", Position: "PCIe-5" },
            { Name: "Riser 2 - PCIe slot 6", Label: "6 PCIe", Position: "PCIe-6" },
            { Name: "Riser 3 - PCIe slot 7", Label: "7 PCIe", Position: "PCIe-7" },
            { Name: "Riser 3 - PCIe slot 8", Label: "8 PCIe", Position: "PCIe-8" },
            { Name: "Riser 4 - PCIe slot 9", Label: "9 PCIe", Position: "PCIe-9" },
            { Name: "Riser 4 - PCIe slot 10", Label: "10 PCIe", Position: "PCIe-10" },
            { Name: "Front Riser 1 - PCIe slot 11", Label: "11 PCIe", Position: "PCIe-11" },
            { Name: "Front Riser 1 - PCIe slot 12", Label: "12 PCIe", Position: "PCIe-12" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "ThinkSystem SR850 V2": {
        Manufacturer: "Lenovo",
        Model: "ThinkSystem SR850 V2",
        Slug: "lenovo-thinksystem-sr850-v2",
        UHeight: 2,
        PartNumber: "7D32",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 40,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU 1", Label: "", Position: "PSU1" },
            { Name: "PSU 2", Label: "", Position: "PSU2" },
            { Name: "PCIe1", Label: "", Position: "Slot1" },
            { Name: "PCIe2", Label: "", Position: "Slot2" },
            { Name: "PCIe3", Label: "", Position: "Slot3" },
            { Name: "OCP3", Label: "", Position: "Slot4" },
            { Name: "PCIe5", Label: "", Position: "Slot5" },
            { Name: "PCIe6", Label: "", Position: "Slot6" },
            { Name: "PCIe7", Label: "", Position: "Slot7" },
            { Name: "PCIe8", Label: "", Position: "Slot8" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "XCC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
