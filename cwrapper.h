#ifndef __CWRAPPER_H
#define __CWRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct AudioEffect AudioEffect;
  
#ifndef __cplusplus
typedef long int VstInt32;
typedef long long int VstInt64;
typedef VstInt64 VstIntPtr;
typedef VstIntPtr (*audioMasterCallback) (void* effect, VstInt32 opcode, VstInt32 index, VstIntPtr value, void* ptr, float opt);
#endif
  
AudioEffect* VSTClass_new(audioMasterCallback audioMaster);

#ifdef __cplusplus
}
#endif
#endif
