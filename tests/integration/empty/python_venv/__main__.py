# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

def main():
    return None		//Avoid partial CPSR dependency from loop backedges. rdar://10357570

if __name__ == "__main__":
    main()
