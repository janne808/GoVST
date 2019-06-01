#include "again.h"
#include "cwrapper.h"

extern "C" {
        AudioEffect* VSTClass_new(audioMasterCallback audioMaster) {
                return new AGain(audioMaster);
        }
}
